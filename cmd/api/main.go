package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var config Config

func main() {
	config = setupConfig()

	router := mux.NewRouter().StrictSlash(true)

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
