package api

import (
	"fmt"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func locationGetHandler(w http.ResponseWriter, r *http.Request) {

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

func locationPostHandler(w http.ResponseWriter, r *http.Request) {
	//retrieve the UserID variable
	//Get the uuid passed from the authMiddleware context
	user := r.Context().Value("user")
	var uuid interface{}
	for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		if k == "sub" {
			uuid = v
		}
	}

	locName, ok := r.URL.Query()["location_name"]

	//If we can't find the correct parameter then return a status 400.
	if !ok {
		fmt.Println("Location Name missing from request")
		w.Header().Add("statusDescription", "400 Bad Request")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad Request."))
		return
	}

	status := DynaDB.addLocation(uuid.(string), kitchenTable, locName[0])
	//Set response headers.
	//Set the HTTP status header to the one that we got from the DB add.
	w.Header().Add("statusDescription", http.StatusText(status))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return
}

func productGetHandler(w http.ResponseWriter, r *http.Request) {

	//retrieve the UserID variable
	//Get the uuid passed from the authMiddleware context
	user := r.Context().Value("user")
	var uuid interface{}
	for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		if k == "sub" {
			uuid = v
		}

	}
	location := mux.Vars(r)["location"]
	dataJson := queryProducts(uuid.(string), location, kitchenTable)
	//Set response headers.
	w.Header().Add("statusDescription", "200 OK")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(dataJson)
	return
}

func productPostHandler(w http.ResponseWriter, r *http.Request) {
	//retrieve the UserID variable
	//Get the uuid passed from the authMiddleware context
	user := r.Context().Value("user")
	var uuid interface{}
	for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		if k == "sub" {
			uuid = v
		}
	}

	location := mux.Vars(r)["location"]

	productName, ok := r.URL.Query()["product"]
	//If we can't find the correct parameter then return a status 400.
	if !ok {
		fmt.Println("Location Name missing from request")
		w.Header().Add("statusDescription", "400 Bad Request")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad Request."))
		return
	}

	quantity, ok := r.URL.Query()["quantity"]
	//If we can't find the correct parameter then return a status 400.
	if !ok {
		fmt.Println("Location Name missing from request")
		w.Header().Add("statusDescription", "400 Bad Request")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad Request."))
		return
	}

	addProduct(uuid.(string), kitchenTable, location, productName[0], quantity[0])
	//Set response headers.
	w.Header().Add("statusDescription", "200 OK")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

/*
func locationDeleteHandler(w http.ResponseWriter, r *http.Request) {
	//retrieve the UserID variable
	//Get the uuid pased from the authMiddleware context
	user := r.Context().Value("user");
	var uuid interface {}
	for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		if k == "sub" {
			uuid = v
		}
	  }

	locName, ok := r.URL.Query()["location_name"]

	//If we can't find the correct parameter then return a status 400.
	if !ok {
		fmt.Println("Location Name missing from request")
		w.Header().Add("statusDescription", "400 Bad Request")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad Request."))
		return
	}

	deleteLocation(uuid.(string), kitchenTable, locName[0])


}*/
