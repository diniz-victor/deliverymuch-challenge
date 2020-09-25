package recipes

import (
	"sort"
	"strings"

	"github.com/deliverymuch-challenge/services/recipes/apigiphy"
	"github.com/deliverymuch-challenge/services/recipes/apipuppy"
)

//Process represents the main process
func Process(parameters string, parametersSlice []string) recipesResponse {
	rp := apipuppy.GetRecipePuppy(parameters)

	sort.Slice(rp.Results, func(i, j int) bool {
		return rp.Results[i].Title < rp.Results[j].Title
	})

	return buildRecipesResponse(rp, parametersSlice)
}

func buildRecipesResponse(pr apipuppy.RecipePuppyResponse, s []string) recipesResponse {
	var r []recipesResponseData
	for _, puppyResponse := range pr.Results {
		responseData := recipesResponseData{
			Title:       puppyResponse.Title,
			Ingredients: strings.Split(puppyResponse.Ingredients, ","),
			Link:        puppyResponse.Href,
			Gif:         apigiphy.GetGiphyLink(puppyResponse.Title),
		}

		r = append(r, responseData)
	}

	return recipesResponse{
		Keywords: s,
		Recipes:  r,
	}
}
