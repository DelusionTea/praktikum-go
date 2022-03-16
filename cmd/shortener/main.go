package main

import (
	"context"
	"github.com/DelusionTea/praktikum-go/cmd/conf"
	"github.com/DelusionTea/praktikum-go/internal/app/handlers"
	"github.com/DelusionTea/praktikum-go/internal/memory"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func setupRouter(repo memory.MemoryInterface, baseURL string) *gin.Engine {
	router := gin.Default()

	handler := handlers.New(repo, baseURL)

	router.GET("/:id", handler.HandlerGetURLByID)
	router.POST("/", handler.HandlerCreateShortURL)
	router.POST("/api/shorten", handler.HandlerShortenURL)

	router.HandleMethodNotAllowed = true

	return router
}

func main() {
	cfg := conf.GetConfig()

	handler := setupRouter(memory.NewMemoryFile(cfg.FilePath), cfg.BaseURL)

	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: handler,
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		log.Fatal(server.ListenAndServe())
		cancel()
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	select {
	case <-sigint:
		cancel()
	case <-ctx.Done():
	}
	server.Shutdown(context.Background())
}
