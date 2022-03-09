package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/DelusionTea/praktikum-go/internal/app/handlers"
	"log"
	"net/http"
)

func main() {

	router := httprouter.New()
	router.POST("/", handlers.HandlerCreateShortURL)
	router.GET("/:id", handlers.HandlerGetURLByID)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
