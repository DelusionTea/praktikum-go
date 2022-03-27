package handlers

import (
	"encoding/json"
	"github.com/DelusionTea/praktikum-go/internal/DataBase"
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

func New(repo memory.MemoryMap, baseURL string) *Handler {
	return &Handler{
		repo:    repo,
		baseURL: baseURL,
	}
}

func (h *Handler) HandlerGetURLByID(c *gin.Context) {
	result := map[string]string{}
	//long, err := h.repo.GetURL(c.Param("id"))
	long, err := shorter.GetURL(c.Param("id"), h.repo)
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
	short := shorter.AddURL(string(body), h.repo)
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
	short := shorter.AddURL(url.URL, h.repo)
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
