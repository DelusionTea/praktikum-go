package handlers

import (
	"context"
	"encoding/json"
	"github.com/DelusionTea/praktikum-go/internal/DataBase"
	"errors"
	"github.com/DelusionTea/praktikum-go/internal/app/shorter"
	"github.com/DelusionTea/praktikum-go/internal/memory"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type PostURL struct {
	URL string
}

type Handler struct {
	repo    memory.MemoryMap
	baseURL string
}
type ManyPostURL struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type ManyPostResponse struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

type ResponseGetURL struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func AddURL(longURL string, repo memory.MemoryMap, user string) string {
	log.Println("Start Add URL")
	shortURL := shorter.Shorter(longURL)
	repo.Values[shortURL] = longURL
	repo.UsersURL[user] = append(repo.UsersURL[user], shortURL)
	repo.WriteRow(longURL, shortURL, repo.FilePath, user)
	log.Println("End Add URL :")
	log.Print(shortURL)
	return shortURL
}

//func AddURLbyID(longURL string, repo memory.MemoryMap, user string) string {
//	log.Println("Start Add URL")
//	shortURL := shorter.Shorter(longURL)
//	repo.Values[shortURL] = longURL
//	repo.UsersURL[user] = append(repo.UsersURL[user], shortURL)
//	repo.WriteRow(longURL, shortURL, repo.FilePath, user)
//	log.Println("End Add URL :")
//	log.Print(shortURL)
//	return shortURL
//}

func GetURL(shortURL string, repo memory.MemoryMap) (string, error) {
	log.Println("Start Get URL")
	resultURL, okey := repo.Values[shortURL]
	log.Println("End Get URL :")
	log.Print(resultURL)
	if !okey {
		return "", errors.New("not found")
	}
	return resultURL, nil
}

func GetUserURL(ctx context.Context, user string, repo memory.MemoryMap) ([]ResponseGetURL, error) {
	result := []ResponseGetURL{}
	for _, url := range repo.UsersURL[user] {
		temp := ResponseGetURL{
			ShortURL:    repo.BaseURL + url,
			OriginalURL: repo.Values[url],
		}
		result = append(result, temp)
	}

	return result, nil
}

func New(repo memory.MemoryMap, baseURL string) *Handler {
	return &Handler{
		repo:    repo,
		baseURL: baseURL,
	}
}

type ShortnerInterface interface {
	//AddURL(longURL string, repo memory.MemoryMap) string
	//GetURL(shortURL string, repo memory.MemoryMap) (string, error)
	AddURL(ctx context.Context, longURL string, shortURL string, user string) error
	GetURL(ctx context.Context, shortURL string) (string, error)
	GetUserURL(ctx context.Context, user string) ([]ResponseGetURL, error)
	AddManyURL(ctx context.Context, urls []ManyPostURL, user string) ([]ManyPostResponse, error)
}

func (h *Handler) HandlerGetURLByID(c *gin.Context) {
	result := map[string]string{}
	//long, err := h.repo.GetURL(c.Param("id"))
	long, err := GetURL(c.Param("id"), h.repo)
	//short := shorter.AddURL(string(body), h.repo)

	if err != nil {
		result["detail"] = err.Error()
		c.IndentedJSON(http.StatusNotFound, result)
		return
	}

	c.Header("Location", long)
	c.String(http.StatusTemporaryRedirect, "")
}

func (h *Handler) HandlerCreateShortURL(c *gin.Context) {

	result := map[string]string{}
	defer c.Request.Body.Close()

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		result["detail"] = "Bad request"
		c.IndentedJSON(http.StatusBadRequest, result)
		return
	}
	short := AddURL(string(body), h.repo, c.GetString("userId"))
	//short := h.repo.AddURL(string(body))
	c.String(http.StatusCreated, h.baseURL+short)
}

func (h *Handler) HandlerShortenURL(c *gin.Context) {
	log.Println("start Shorten")
	headerContentType := c.GetHeader("Content-Type")
	if headerContentType != "application/json" {
		c.IndentedJSON(http.StatusUnsupportedMediaType, headerContentType)
		return
	}
	result := map[string]string{}
	var url PostURL
	defer c.Request.Body.Close()
	log.Println("Start read. Body:  ")
	body, err := ioutil.ReadAll(c.Request.Body)
	log.Print(body)
	if err != nil {
		result["detail"] = "Bad request"
		c.IndentedJSON(http.StatusBadRequest, result)
		return
	}
	json.Unmarshal(body, &url)
	if url.URL == "" {
		result["detail"] = "Bad request"
		c.IndentedJSON(http.StatusBadRequest, result)
		return
	}
	short := AddURL(url.URL, h.repo, c.GetString("userId"))
	//short := h.repo.AddURL(url.URL)
	result["result"] = h.baseURL + short
	c.IndentedJSON(http.StatusCreated, result)

}

func (h *Handler) HandlerPingDB(c *gin.Context) {
	//При успешной проверке хендлер должен вернуть HTTP-статус 200 OK, при неуспешной — 500 Internal Server Error.
	//err := DataBase.Ping(c.Request.Context()
	ctx := c.Request.Context()
	err := DataBase.PGDataBase.Ping(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, "")
		return
	}
	c.String(http.StatusOK, "")
}

func (h *Handler) HandlerHistoryOfURLs(c *gin.Context) {
	//result, err := h.repo.GetUserURL(c.Request.Context(), c.GetString("userId"))
	log.Println("start HandlerHistoryOfURLs")
	log.Println(c.GetString("userId"))
	//log.Println(c.GetString(c.Request.Cookie("userId")))
	result, err := GetUserURL(c.Request.Context(), c.GetString("userId"), h.repo)
	log.Println(result)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Println("StatusInternalServerError")
		return
	}
	if len(result) == 0 {
		c.IndentedJSON(http.StatusNoContent, result)
		log.Println("StatusNoContent")
		return
	}
	log.Println("StatusOK")
	log.Println(result)
	c.IndentedJSON(http.StatusOK, result)
}
