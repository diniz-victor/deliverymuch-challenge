package recipes

//RecipesResponseData represents the Recipes inside the main response
type RecipesResponseData struct {
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Link        string   `json:"link"`
	Gif         string   `json:"gif"`
}

//RecipesResponse represents the main response
type RecipesResponse struct {
	Keywords []string              `json:"keywords"`
	Recipes  []RecipesResponseData `json:"recipes"`
}
