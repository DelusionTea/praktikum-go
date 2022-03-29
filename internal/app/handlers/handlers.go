package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DelusionTea/praktikum-go/internal/app/shorter"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type DBError struct {
	Err   error
	Title string
}

func (err *DBError) Error() string {
	return fmt.Sprintf("%v", err.Err)
}

func (err *DBError) Unwrap() error {
	return err.Err
}

func NewErrorWithDB(err error, title string) error {
	return &DBError{
		Err:   err,
		Title: title,
	}
}

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

func (h *Handler) handleError(c *gin.Context, err error) {
	message := make(map[string]string)
	message["detail"] = err.Error()
	c.IndentedJSON(http.StatusBadRequest, message)
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
		var ue *DBError
		if errors.As(err, &ue) && ue.Title == "Deleted" {
			c.Status(http.StatusGone)
			return
		}
		result["detail"] = err.Error()
		c.IndentedJSON(http.StatusNotFound, result)
		return
	}

	c.Header("Location", long)
	c.String(http.StatusTemporaryRedirect, "")
}

func (h *Handler) HandlerCreateShortURL(c *gin.Context) {

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
		var ue *DBError
		if errors.As(err, &ue) && ue.Title == "UniqConstraint" {
			c.String(http.StatusConflict, h.baseURL+shortURL)
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.String(http.StatusCreated, h.baseURL+shortURL)
}

func (h *Handler) HandlerShortenURL(c *gin.Context) {
	result := map[string]string{}
	var url PostURL

	defer c.Request.Body.Close()

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		h.handleError(c, err)
		return
	}
	json.Unmarshal(body, &url)
	if url.URL == "" {
		h.handleError(c, errors.New("bad request"))
		return
	}
	shortURL := shorter.Shorter(url.URL)
	err = h.repo.AddURL(c.Request.Context(), url.URL, shortURL, c.GetString("userId"))
	if err != nil {
		var ue *DBError
		if errors.As(err, &ue) && ue.Title == "UniqConstraint" {
			result["result"] = h.baseURL + shortURL
			c.IndentedJSON(http.StatusConflict, result)
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	result["result"] = h.baseURL + shortURL
	c.IndentedJSON(http.StatusCreated, result)
}

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
