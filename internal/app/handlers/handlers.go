package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DelusionTea/praktikum-go/internal/app/shorter"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type PostURL struct {
	URL string
}

type Handler struct {
	repo    ShorterInterface
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

//func AddURL(longURL string, repo ShorterInterface, user string) string {
//	log.Println("Start Add URL")
//	shortURL := shorter.Shorter(longURL)
//	repo.Values[shortURL] = longURL
//	repo.UsersURL[user] = append(repo.UsersURL[user], shortURL)
//	repo.WriteRow(longURL, shortURL, repo.FilePath, user)
//	log.Println("End Add URL :")
//	log.Print(shortURL)
//	return shortURL
//}

//func AddURLbyID(longURL string, repo ShorterInterface, user string) string {
//	log.Println("Start Add URL")
//	shortURL := shorter.Shorter(longURL)
//	repo.Values[shortURL] = longURL
//	repo.UsersURL[user] = append(repo.UsersURL[user], shortURL)
//	repo.WriteRow(longURL, shortURL, repo.FilePath, user)
//	log.Println("End Add URL :")
//	log.Print(shortURL)
//	return shortURL
//}

//func GetURL(shortURL string, repo ShorterInterface) (string, error) {
//	log.Println("Start Get URL")
//	resultURL, okey := repo.Values[shortURL]
//	log.Println("End Get URL :")
//	log.Print(resultURL)
//	if !okey {
//		return "", errors.New("not found")
//	}
//	return resultURL, nil
//}

//func GetUserURL(ctx context.Context, user string, repo ShorterInterface) ([]ResponseGetURL, error) {
//	result := []ResponseGetURL{}
//	for _, url := range repo.UsersURL[user] {
//		temp := ResponseGetURL{
//			ShortURL:    repo.BaseURL + url,
//			OriginalURL: repo.Values[url],
//		}
//		result = append(result, temp)
//	}
//
//	return result, nil
//}

func New(repo ShorterInterface, baseURL string) *Handler {
	return &Handler{
		repo:    repo,
		baseURL: baseURL,
	}
}

type ShorterInterface interface {
	//AddURL(longURL string, repo ShorterInterface) string
	//GetURL(shortURL string, repo ShorterInterface) (string, error)
	AddURL(ctx context.Context, longURL string, shortURL string, user string) error
	GetURL(ctx context.Context, shortURL string) (string, error)
	GetUserURL(ctx context.Context, user string) ([]ResponseGetURL, error)
	AddManyURL(ctx context.Context, urls []ManyPostURL, user string) ([]ManyPostResponse, error)
	Ping(ctx context.Context) error
}

func (h *Handler) HandlerBatch(c *gin.Context) {
	var data []ManyPostURL
	defer c.Request.Body.Close()

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, c.GetString(""))
		return
	}
	json.Unmarshal(body, &data)
	fmt.Println(data)
	response, err := h.repo.AddManyURL(c.Request.Context(), data, c.GetString("userId"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, c.GetString("userId"))
		return
	}
	if response == nil {
		c.IndentedJSON(http.StatusBadRequest, "Error")
		return
	}
	c.IndentedJSON(http.StatusCreated, response)
}
func (h *Handler) HandlerGetURLByID(c *gin.Context) {
	result := map[string]string{}
	//long, err := h.repo.GetURL(c.Param("id"))
	long, err := h.repo.GetURL(c, c.Param("id"))
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

	//result := map[string]string{}
	//defer c.Request.Body.Close()
	//
	//body, err := ioutil.ReadAll(c.Request.Body)
	//
	//if err != nil {
	//	result["detail"] = "Bad request"
	//	c.IndentedJSON(http.StatusBadRequest, result)
	//	return
	//}
	//short := h.repo.AddURL(string(body), h.repo, c.GetString("userId"))
	////short := h.repo.AddURL(string(body))
	//c.String(http.StatusCreated, h.baseURL+short)

	defer c.Request.Body.Close()

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Error")
		return
	}
	longURL := string(body)
	shortURL := shorter.Shorter(longURL)
	err = h.repo.AddURL(c.Request.Context(), longURL, shortURL, c.GetString("userId"))
	if err != nil {
		//var ue *ErrorWithDB
		//if errors.As(err, &ue) && ue.Title == "UniqConstraint" {
		//	c.String(http.StatusConflict, h.baseURL+shortURL)
		//	return
		//}
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.String(http.StatusCreated, h.baseURL+shortURL)
}

func (h *Handler) HandlerShortenURL(c *gin.Context) {
	//log.Println("start Shorten")
	//headerContentType := c.GetHeader("Content-Type")
	//if headerContentType != "application/json" {
	//	c.IndentedJSON(http.StatusUnsupportedMediaType, headerContentType)
	//	return
	//}
	//result := map[string]string{}
	//var url PostURL
	//defer c.Request.Body.Close()
	//log.Println("Start read. Body:  ")
	//body, err := ioutil.ReadAll(c.Request.Body)
	//log.Print(body)
	//if err != nil {
	//	result["detail"] = "Bad request"
	//	c.IndentedJSON(http.StatusBadRequest, result)
	//	return
	//}
	//json.Unmarshal(body, &url)
	//if url.URL == "" {
	//	result["detail"] = "Bad request"
	//	c.IndentedJSON(http.StatusBadRequest, result)
	//	return
	//}
	//short := h.repo.AddURL(url.URL, h.repo, c.GetString("userId"))
	////short := h.repo.AddURL(url.URL)
	//result["result"] = h.baseURL + short
	//c.IndentedJSON(http.StatusCreated, result)
	result := map[string]string{}
	var url PostURL

	defer c.Request.Body.Close()

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, result)
		return
	}
	json.Unmarshal(body, &url)
	if url.URL == "" {
		c.IndentedJSON(http.StatusBadRequest, result)
		return
	}
	shortURL := shorter.Shorter(url.URL)
	err = h.repo.AddURL(c.Request.Context(), url.URL, shortURL, c.GetString("userId"))
	if err != nil {
		//var ue *ErrorWithDB
		//if errors.As(err, &ue) && ue.Title == "UniqConstraint" {
		//	result["result"] = h.baseURL + shortURL
		//	c.IndentedJSON(http.StatusConflict, result)
		//	return
		//}
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	result["result"] = h.baseURL + shortURL
	c.IndentedJSON(http.StatusCreated, result)
}

//func (h *Handler) HandlerPingDB(c *gin.Context) {
//	//При успешной проверке хендлер должен вернуть HTTP-статус 200 OK, при неуспешной — 500 Internal Server Error.
//	//err := DataBase.Ping(c.Request.Context()
//	ctx := c.Request.Context()
//	err := DataBase.PGDataBase.Ping(ctx)
//	if err != nil {
//		c.String(http.StatusInternalServerError, "")
//		return
//	}
//	c.String(http.StatusOK, "")
//}

func (h *Handler) HandlerPingDB(c *gin.Context) {
	err := h.repo.Ping(c.Request.Context())
	if err != nil {
		c.String(http.StatusInternalServerError, "")
		return
	}
	c.String(http.StatusOK, "")
}

func (h *Handler) HandlerHistoryOfURLs(c *gin.Context) {
	result, err := h.repo.GetUserURL(c.Request.Context(), c.GetString("userId"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	if len(result) == 0 {
		c.IndentedJSON(http.StatusNoContent, result)
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
