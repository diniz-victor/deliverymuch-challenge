package recipes

//GiphyGateway represents gateway to giphy
type GiphyGateway interface {
	GetGiphy(query string) string
}

//PuppyGateway represents gateway to puppy
type PuppyGateway interface {
	GetPuppy(parameters string) RecipePuppyResponse
}

//Service represents gateway to service
type Service interface {
	Process(parameters string, parametersSlice []string) RecipesResponse
}
