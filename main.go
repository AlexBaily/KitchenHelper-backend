package main

import (
	"log"
	"net/http"
	"os"

	"github.com\alexbaily\api\routes" routes
)

//"Global" variables.
var (
	kitchenTable   string = os.Getenv("KITCHTABLE")
	tokenVerifyURL string = os.Getenv("TOKENURL")
)



func main() {
	//Create a new mux router.
	router := routes.setRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
