package api

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

func locationHandler(w http.ResponseWriter, r *http.Request) {
	//Set response headers.
	w.Header().Add("statusDescription", "200 OK")
	w.Header().Set("statusDescription", "200 OK")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//retrieve the UserID variable

	//Get the uuid pased from the authMiddleware context
	
	user := r.Context().Value("user");
	var uuid interface {}
	for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		if k == "sub" {
			uuid = v
		}
		
	  }
	dataJson := queryLocations(uuid.(string), kitchenTable)
	w.Write(dataJson)
}
