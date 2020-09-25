package recipes

//RecipesResponseData represents the Recipes inside the main response
type recipesResponseData struct {
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Link        string   `json:"link"`
	Gif         string   `json:"gif"`
}

//RecipesResponse represents the main response
type recipesResponse struct {
	Keywords []string              `json:"keywords"`
	Recipes  []recipesResponseData `json:"recipes"`
}
