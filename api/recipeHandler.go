package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alexbaily/KitchenHelper-backend/models"
	"github.com/gorilla/mux"
)

func recipeGetHandler(w http.ResponseWriter, r *http.Request) {
	//retrieve the UserID variable
	uuid := getUserID(r.Context().Value("user"))

	recipe := mux.Vars(r)["recipe"]

	dataJson := DynaDB.queryRecipes(uuid, recipe, recipeTable)
	//Set response headers.
	w.Header().Add("statusDescription", "200 OK")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(dataJson)
	return
}

/*
	recipePostHandler will take a JSON document, parse and then add to the DB.
	JSON syntax:
	{
		Name,
		Description,
		Ingredients : [
			{
				Name,
				Quantity,
				Measurement
			}
		],
		Steps: [
			{
				ID,
				Text,
				ImgURL
			}
		],
		Notes,
		Sharing
	}
*/
func recipePostHandler(w http.ResponseWriter, r *http.Request) {
	//retrieve the UserID variable
	uuid := getUserID(r.Context().Value("user"))

	recipeJson, err := ioutil.ReadAll(r.Body)
	//If we can't find the correct parameter then return a status 400.
	if err != nil {
		fmt.Println("Error reading from the request body.")
		w.Header().Add("statusDescription", "500 Internal Server Error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
		return
	}

	recipe := models.RecipeRecord{}
	err = json.Unmarshal(recipeJson, &recipe)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
	}
	fmt.Printf("%v", recipe)

	addRecipe(uuid, recipe, recipeTable)
	//Set response headers.
	w.Header().Add("statusDescription", "200 OK")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}
