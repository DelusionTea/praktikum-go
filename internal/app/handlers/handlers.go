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
	repo    memory.MemoryMap
	baseURL string
	result  BodyResponse
}

const endpoint = "http://localhost:8080/"

func NewHandler(config *conf.Config) *Handler {
	return &Handler{
		baseURL: config.BaseURL,
		repo:    *memory.NewMemoryMap(config.FilePath),
	}

}
func (h *Handler) CallHandlers(router chi.Router) {
	router.Post("/", h.HandlerCreateShortURL)
	router.Get("/{ID}", h.HandlerGetURLByID)
	router.Post("/api/shorten", h.HandlerShortenURL)

}

func (h *Handler) HandlerCreateShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return
	}
	//long := string(body)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	//short := shorter.shorter(h.baseURL)
	short := h.repo.AddURL(string(body))
	w.Write([]byte(short))
}

func (h *Handler) HandlerGetURLByID(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "ID")
	log.Println(param)
	param = h.baseURL + param
	long, err := h.repo.GetURL(param)
	if err != nil {
		//http.Error(w, "Error", http.StatusBadRequest)
		//return
	}
	log.Println(long)
	if long == "" {
		http.Error(w, "id error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Location", long)
	w.WriteHeader(http.StatusTemporaryRedirect)

}

func (h *Handler) HandlerShortenURL(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&h.baseURL); err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		//return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	shortURL := h.repo.AddURL(h.baseURL)
	h.result.ResultURL = shortURL
	result, err := json.Marshal(h.result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(result))
}
