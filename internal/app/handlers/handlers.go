package handlers

import (
	"encoding/json"
	"github.com/DelusionTea/praktikum-go/internal/memory"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type PostURL struct {
	URL string
}

type Handler struct {
	repo    memory.MemoryInterface
	baseURL string
}

func New(repo memory.MemoryInterface, baseURL string) *Handler {
	return &Handler{
		repo:    repo,
		baseURL: baseURL,
	}
}

func (h *Handler) HandlerGetURLByID(c *gin.Context) {
	result := map[string]string{}
	long, err := h.repo.GetURL(c.Param("id"))

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
	short := h.repo.AddURL(string(body))
	c.String(http.StatusCreated, h.baseURL+short)
}

func (h *Handler) HandlerShortenURL(c *gin.Context) {
	result := map[string]string{}
	var url PostURL

	defer c.Request.Body.Close()

	body, err := ioutil.ReadAll(c.Request.Body)

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

	short := h.repo.AddURL(url.URL)
	result["result"] = h.baseURL + short
	c.IndentedJSON(http.StatusCreated, result)

}
