package api

import "github.com/gorilla/mux"

func setRoutes() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/exercises", locationHandler)
	router.Use(authMiddleware)
	return router
}
