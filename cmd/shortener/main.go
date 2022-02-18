package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/DelusionTea/praktikum-go/internal/app/handlers"
	"log"
	"net/http"
	Shorting "../../internal/app/handlers/urlshorter.go"
)

func main() {
	//http.HandleFunc("/", myrequest)
	router := httprouter.New()
	router.POST("/", Shorting.HandlerCreateShortURL)
	router.GET("/:id", Shorting.HandlerGetURLByID)
	//Сервер должен быть доступен по адресу: http://localhost:8080.
	//http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
