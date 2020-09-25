package apipuppy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//GetRecipePuppy get the recipe from puppy endpoint
func GetRecipePuppy(p string) RecipePuppyResponse {
	resp, err := http.Get(fmt.Sprintf("http://www.recipepuppy.com/api/?i=%s", p))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var recipePuppyResponse RecipePuppyResponse
	err = json.Unmarshal(bodyBytes, &recipePuppyResponse)
	if err != nil {
		log.Fatalln(err)
	}

	return recipePuppyResponse
}
