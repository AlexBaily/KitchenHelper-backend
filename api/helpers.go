package api

import (
	"encoding/json"
	"fmt"

	"github.com/alexbaily/KitchenHelper-backend/models"
	"github.com/dgrijalva/jwt-go"
)

func getUserID(user interface{}) string {
	//retrieve the UserID variable
	//Get the uuid passed from the authMiddleware context
	var uuid interface{}
	for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		if k == "sub" {
			uuid = v
		}

	}
	return uuid.(string)
}

//Takes a recipe json
//turns into a recipe record
func recipeFromJson(jsonRecipe []byte) models.RecipeRecord {
	var record models.RecipeRecord
	if err := json.Unmarshal(jsonRecipe, &record); err != nil {
		fmt.Println(err)
		return nil
	}

	return record
}
