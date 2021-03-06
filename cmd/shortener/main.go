package main

import (
	"context"
	"database/sql"
	"github.com/DelusionTea/praktikum-go/cmd/conf"
	"github.com/DelusionTea/praktikum-go/internal/DataBase"
	"github.com/DelusionTea/praktikum-go/internal/app/handlers"
	"github.com/DelusionTea/praktikum-go/internal/app/middleware"
	"github.com/DelusionTea/praktikum-go/internal/memory"
	"github.com/DelusionTea/praktikum-go/internal/workers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func setupRouter(repo handlers.ShorterInterface, conf *conf.Config, wp *workers.Workers) *gin.Engine {
	/*func setupRouter(repo memory.MemoryMap, baseURL string, conf *conf.Config) *gin.Engine {*/
	router := gin.Default()
	//router.
	router.Use(middleware.GzipEncodeMiddleware())
	router.Use(middleware.GzipDecodeMiddleware())
	router.Use(middleware.CookieMiddleware(conf))
	//router.Use(gzip.Gzip(gzip.DefaultCompression))
	handler := handlers.New(repo, conf.BaseURL, wp)

	router.GET("/:id", handler.HandlerGetURLByID)
	router.POST("/", handler.HandlerCreateShortURL)
	router.POST("/api/shorten", handler.HandlerShortenURL)
	router.GET("/ping", handler.HandlerPingDB)
	router.GET("/api/user/urls", handler.HandlerHistoryOfURLs)
	//POST /api/shorten/batch
	router.POST("/api/shorten/batch", handler.HandlerBatch)
	router.DELETE("/api/user/urls", handler.DeleteBatch)

	router.HandleMethodNotAllowed = true

	return router
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cfg := conf.GetConfig()
	var handler *gin.Engine
	//db, err := sql.Open("postgres", cfg.DataBase)
	wp := workers.New(ctx, cfg.NumbWorkers, cfg.WorkerBuff)
	go func() {
		wp.Run(ctx)
	}()
	if cfg.DataBase != "" {
		//handler = setupRouter(DataBase.NewDatabase(cfg.BaseURL, cfg.DataBase))
		db, err := sql.Open("postgres", cfg.DataBase)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		database.SetUpDataBase(db, ctx)
		handler = setupRouter(database.NewDatabaseRepository(cfg.BaseURL, db), cfg, wp)
		//handler = setupRouter(memory.NewMemoryFile(cfg.FilePath, cfg.BaseURL), cfg.BaseURL, cfg)
	} else {
		handler = setupRouter(memory.NewMemoryFile(ctx, cfg.FilePath, cfg.BaseURL), cfg, wp)
	}
	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: handler,
	}

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
