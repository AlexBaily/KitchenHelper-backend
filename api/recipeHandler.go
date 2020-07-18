package api

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func recipeGetHandler(w http.ResponseWriter, r *http.Request) {

	//retrieve the UserID variable
	//Get the uuid passed from the authMiddleware context
	user := r.Context().Value("user")
	var uuid interface{}
	for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		if k == "sub" {
			uuid = v
		}

	}
	dataJson := DynaDB.queryLocations(uuid.(string), kitchenTable)
	//Set response headers.
	w.Header().Add("statusDescription", "200 OK")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(dataJson)
	return
}
