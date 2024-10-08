package models

// Ingredients are an array of strings on the Food API endpoint
type Ingredients struct {
	IngredientsArray []string `json:"ingredients"`
}
