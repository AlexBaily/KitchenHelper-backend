package main

import (
	"os"

	"github.com/alexbaily/KitchenHelper-backend/api"
)

var (
	runInLambda string = os.Getenv("LAMBDA")
)

func main() {
	if runInLambda == "lambda" {
		api.InitLambda()
	} else {
		api.InitServer()
	}
}
