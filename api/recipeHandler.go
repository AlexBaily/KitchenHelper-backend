package api

import (
	"net/http"
)

func recipeGetHandler(w http.ResponseWriter, r *http.Request) {
	//retrieve the UserID variable
	uuid := getUserID(r.Context().Value("user"))

	dataJson := DynaDB.queryRecipes(uuid, recipeTable)
	//Set response headers.
	w.Header().Add("statusDescription", "200 OK")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(dataJson)
	return
}
