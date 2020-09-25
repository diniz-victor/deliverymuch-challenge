package errors

//Response represents the json response error
type Response struct {
	Error string `json:"error"`
}

//ErrorIngredientsExceededLimit error message for more than 3 ingredients
var ErrorIngredientsExceededLimit = "Please try 3 ingredients at maximum."
