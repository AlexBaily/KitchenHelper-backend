package api

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//Setup the routes and middlware
func setRoutes() (n *negroni.Negroni) {

	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/locations", locationGetHandler).
		Methods("GET")
	//Unsure whether to use POST or PUT here.
	//Using POST as it's /locations not /locations/{id}
	router.HandleFunc("/locations", locationPostHandler).
		Methods("POST")
	router.HandleFunc("/locations/{location}", productGetHandler).
		Methods("GET")
	router.HandleFunc("/locations/{location}", productPostHandler).
		Methods("POST")
	router.HandleFunc("/recipes", recipeGetHandler).
		Methods("GET")
	router.HandleFunc("/recipes/{recipe}", recipeGetHandler).
		Methods("GET")
	router.HandleFunc("/recipes/{recipe}", recipePostHandler).
		Methods("POST")
	//Set the jwt handler which will verify the token.
	//Let negroni handle this.
	tokenMiddleware := verifyToken()
	n = negroni.New()
	n.Use(negroni.HandlerFunc(tokenMiddleware.HandlerWithNext))
	n.UseHandler(router)

	return n
}
