package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dtsputri/microservice/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))

	fmt.Println("Menu service port:8888")
	log.Panic(http.ListenAndServe(":8888", router))
}
