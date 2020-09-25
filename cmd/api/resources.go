package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/deliverymuch-challenge/entities/errors"
	"github.com/deliverymuch-challenge/services/recipes"
)

func postFunction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := r.URL.Query()["i"][0]

	s := strings.Split(p, ",")
	if len(s) > 3 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.Response{Error: errors.ErrorIngredientsExceededLimit})
		return
	}

	response := recipes.Process(p, s)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
