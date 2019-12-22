package api

import (
	"net/http"
)

func locationHandler(w http.ResponseWriter, r *http.Request) {
	//Set response headers.
	w.Header().Add("statusDescription", "200 OK")
	w.Header().Set("statusDescription", "200 OK")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//retrieve the UserID variable

	//Get the uuid pased from the authMiddleware context
	uuid := r.Context().Value("sub")
	dataJson := queryLocations(uuid.(string), kitchenTable)
	w.Write(dataJson)
}
