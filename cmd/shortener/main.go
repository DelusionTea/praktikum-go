package main

import (
	"github.com/DelusionTea/praktikum-go/cmd/conf"
	"github.com/DelusionTea/praktikum-go/internal/app/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {

	router := chi.NewRouter()
	log.Println(router)
	config := conf.GetConfig()

	handler := handlers.NewHandler(config)
	log.Println(handler)
	handler.CallHandlers(router)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
