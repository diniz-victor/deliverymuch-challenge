package recipes

import (
	"sort"
	"strings"
)

type service struct {
	puppyGateway PuppyGateway
	giphyGateway GiphyGateway
}

//NewService creates a new service
func NewService(p PuppyGateway, g GiphyGateway) Service {
	return &service{
		puppyGateway: p,
		giphyGateway: g,
	}
}

//Process represents the main process
func (s service) Process(parameters string, parametersSlice []string) RecipesResponse {
	rp := s.puppyGateway.GetPuppy(parameters)

	sort.Slice(rp.Results, func(i, j int) bool {
		return rp.Results[i].Title < rp.Results[j].Title
	})

	return s.buildRecipesResponse(rp, parametersSlice)
}

func (s service) buildRecipesResponse(pr RecipePuppyResponse, keywords []string) RecipesResponse {
	var recipeResponse []RecipesResponseData
	for _, puppyResponse := range pr.Results {
		responseData := RecipesResponseData{
			Title:       puppyResponse.Title,
			Ingredients: strings.Split(puppyResponse.Ingredients, ","),
			Link:        puppyResponse.Href,
			Gif:         s.giphyGateway.GetGiphy(puppyResponse.Title),
		}

		recipeResponse = append(recipeResponse, responseData)
	}

	return RecipesResponse{
		Keywords: keywords,
		Recipes:  recipeResponse,
	}
}
