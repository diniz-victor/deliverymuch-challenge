package main

import "github.com/gorilla/mux"

func setupRouter(router *mux.Router) {
	router.
		Methods("GET").
		Path("/recipes").
		HandlerFunc(postFunction)
}
