package recipes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type puppy struct {
	endpoint string
}

//NewPuppy creates a new puppy
func NewPuppy(endpoint string) PuppyGateway {
	return &puppy{endpoint: endpoint}
}

//Get get the recipe from puppy endpoint
func (p puppy) GetPuppy(parameters string) RecipePuppyResponse {
	resp, err := http.Get(fmt.Sprintf("%s?i=%s", p.endpoint, parameters))
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
