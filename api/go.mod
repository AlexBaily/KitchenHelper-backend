module github.com/alexbaily/KitchenHelper-backend/api

go 1.16

replace github.com/alexbaily/KitchenHelper-backend/models => ../models

require (
    github.com/alexbaily/KitchenHelper-backend/models v0.0.0
	github.com/auth0/go-jwt-middleware v1.0.0
	github.com/aws/aws-lambda-go v1.23.0
	github.com/aws/aws-sdk-go v1.38.40
	github.com/awslabs/aws-lambda-go-api-proxy v0.10.0
	github.com/form3tech-oss/jwt-go v3.2.3+incompatible
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/urfave/negroni v1.0.0
)
