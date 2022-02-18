package main

import (
    "net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/DelusionTea/praktikum-go/blob/main/internal/app/handlers/"
)

func main() {
	//http.HandleFunc("/", myrequest)
	router := httprouter.New()
	router.POST("/", HandlerCreateShortURL)
	router.GET("/:id", HandlerGetURLByID)
	//Сервер должен быть доступен по адресу: http://localhost:8080.
	//http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
