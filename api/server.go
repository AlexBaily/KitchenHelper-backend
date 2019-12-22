package api

import (
	"net/http"
	"os"
)

//"Global" variables.
var (
	kitchenTable   string = os.Getenv("KITCHTABLE")
	tokenVerifyURL string = os.Getenv("TOKENURL")
	AUTH_AUDIENCE string = os.Getenv("AUTH_AUDIENCE")
)

func InitServer() {

	router := setRoutes()
	http.ListenAndServe(":8080", router)

}
