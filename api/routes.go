package api

import "github.com/gorilla/mux"

func setRoutes() r *Router {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/exercises", locationHandler)
	r.Use(authMiddleware)
}
