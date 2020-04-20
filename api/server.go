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

//InitServer gets and sets the routes for the http server before
//it starts listening. It is a function so we can use this or Lambda
//which requires a wrapper.
func InitServer() {

	router := setRoutes()
	http.ListenAndServe(":8080", router)

}

//InitServer gets and sets the routes for the http server before
//it starts listening. It also uses the AWS negroniadapter which
//is required for Lambda.
func InitLambda() {

	router := setRoutes()
	negroniLambda := *negroniadapter.New(router)
	lambda.Start(negroniLambda)
}
