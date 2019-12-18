package api

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func getClaims(tokenString string) (jwt.MapClaims, bool) {
	//Parse the JWT token, no verification going on at the moment
	if token, _ := jwt.Parse(tokenString, nil); token != nil {
		log.Printf("parsed token: %+v", token)
		//Get the claims from the parsed token, && token.Valid will be used.
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			log.Printf("claims = %+v", claims)
			return claims, true
		} else {
			log.Printf("Invalid JWT Token, token = %+v", tokenString)
			return nil, false
		}
	} else {
		return nil, false
	}
}

//Middleware to read the Authorization header for the Cognito JWT token
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Don't bother checking the Auth header if we are just going to root
		if r.URL.Path == "/" {
			next.ServeHTTP(w, r)
		} else {
			//Get the token
			token := r.Header.Get("Authorization")
			//log.Printf("token %+v", token)
			jwtToken := strings.Split(token, " ")
			//Check is the jwtToken contains an actual token
			if len(jwtToken) <= 1 {
				//Return a 403 if no token is found
				http.Error(w, "Forbidden", http.StatusForbidden)
			} else {
				claims, ok := getClaims(jwtToken[1])
				if ok {
					log.Printf("jwtAuth %+v", claims["sub"])
					//Add the context to the next request
					//The sub value is the UUID for the user
					ctx := context.WithValue(r.Context(),
						"sub", claims["sub"])
					next.ServeHTTP(w, r.WithContext(ctx))
				}
			}
		}
	})
}
