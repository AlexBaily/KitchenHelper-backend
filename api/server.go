package api

import (
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	negroniadapter "github.com/awslabs/aws-lambda-go-api-proxy/negroni"
)

//"Global" variables.
var (
	kitchenTable   string = os.Getenv("KITCHTABLE")
	tokenVerifyURL string = os.Getenv("TOKENURL")
	AUTH_AUDIENCE  string = os.Getenv("AUTH_AUDIENCE")
)

func InitServer() {

	router := setRoutes()
	http.ListenAndServe(":8080", router)

}

func InitLambda() {

	router := setRoutes()
	negroniLambda := *negroniadapter.New(router)
	lambda.Start(negroniLambda)
}
