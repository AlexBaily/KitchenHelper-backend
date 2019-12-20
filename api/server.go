package api

import (
	"net/http"
	"os"
)

//"Global" variables.
var (
	kitchenTable   string = os.Getenv("KITCHTABLE")
	tokenVerifyURL string = os.Getenv("TOKENURL")
)

func InitServer() {

	router := setRoutes()
	http.ListenAndServe(":8080", router)

}
