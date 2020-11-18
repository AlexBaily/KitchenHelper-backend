package models

//Record struct that will house the DynamoDB records.
type RecipeRecord struct {
	UserID           string
	RecipeIdentifier string
	RecipeName       string
	PhotoURL         string
	Steps            []string
	Ingredients      []string
}
