package recipes

//RecipePuppyResponse represents the response from consumed recipe puppy endpoint
type RecipePuppyResponse struct {
	Title   string  `json:"title"`
	Version float64 `json:"version"`
	Href    string  `json:"href"`
	Results []struct {
		Title       string `json:"title"`
		Href        string `json:"href"`
		Ingredients string `json:"ingredients"`
		Thumbnail   string `json:"thumbnail"`
	} `json:"results"`
}
