package models

//Record struct that will house the DynamoDB records.
type RecipeRecord struct {
	UserID           string              `json:"userID"`
	RecipeIdentifier string              `json:"recipeIdentifier"`
	RecipeName       string              `json:"recipeName"`
	PhotoURL         string              `json:"photoURL"`
	Description      string              `json:"description"`
	Sharing          string              `json:"sharing"`
	Steps            []map[string]string `json:"steps"`
	Ingredients      []map[string]string `json:"ingredients"`
}
