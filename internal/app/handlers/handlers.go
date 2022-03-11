package handlers

import (
	"encoding/json"
	"github.com/DelusionTea/praktikum-go/cmd/conf"
	"github.com/DelusionTea/praktikum-go/internal/memory"
	"github.com/go-chi/chi/v5"
	//"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

type BodyResponse struct {
	ResultURL string `json:"result"`
}

type Handler struct {
	repo    *memory.MemoryMap
	baseURL string
	result  BodyResponse
}

const (
	createURL  = "/"
	getURL     = "/{articleID}"
	shortenURL = "/api/shorten"
)

func NewHandler(config *conf.Config) *Handler {
	return &Handler{
		baseURL: config.BaseURL,
		repo:    memory.NewMemoryMap(config.FilePath),
	}

}
func (h *Handler) CallHandlers(router chi.Router) {
	log.Println("Start Call Handlers")
	router.Post(createURL, h.HandlerCreateShortURL)
	router.Route(getURL, func(r chi.Router) {
		r.Get("/", h.HandlerGetURLByID)
	})
	router.Post(shortenURL, h.HandlerShortenURL)
	//
	//log.Println("i'm here")
	//log.Println(h)
	//router.Post("/", h.HandlerCreateShortURL)
	//router.Get("/{ID}", h.HandlerGetURLByID)
	//router.Post("/api/shorten", h.HandlerShortenURL)

}

func (h *Handler) HandlerCreateShortURL(w http.ResponseWriter, r *http.Request) {
	log.Println("Start Handler Create Short URL")
	body, err := io.ReadAll(r.Body)
	log.Println("body Handler Create Short URL")
	log.Println(body)
	defer r.Body.Close()
	if err != nil {
		log.Println("error Handler Create Short URL")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	//long := string(body)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	//short := shorter.shorter(h.baseURL)
	short := h.repo.AddURL(string(body))
	long, err := h.repo.GetURL(short)
	if err != nil {
		log.Println("error Handler Create Short URL")
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", long)
	//w.Write([]byte(short))
}

func (h *Handler) HandlerGetURLByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Start Handler Get URL By ID")
	param := chi.URLParam(r, "ID")
	log.Println(param)
	param = h.baseURL + param
	long, err := h.repo.GetURL(param)
	if err != nil {
		log.Println("error Handler Get URL By ID")
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
	log.Println(long)
	if long == "" {
		log.Println("error Handler Get URL By ID")
		http.Error(w, "id error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", long)
	w.WriteHeader(http.StatusTemporaryRedirect)

}

func (h *Handler) HandlerShortenURL(w http.ResponseWriter, r *http.Request) {
	log.Println("Start Handler Shorten URL")
	if err := json.NewDecoder(r.Body).Decode(&h.baseURL); err != nil {
		log.Println("error HandlerShortenURL")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	shortURL := h.repo.AddURL(h.baseURL)
	h.result.ResultURL = shortURL
	result, err := json.Marshal(h.result)
	if err != nil {
		log.Println("error HandlerShortenURL")
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(result))
}
