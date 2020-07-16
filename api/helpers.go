package api

import "github.com/dgrijalva/jwt-go"

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
