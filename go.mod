module github.com/alexbaily/KitchenHelper-backend

go 1.16

replace github.com/alexbaily/KitchenHelper-backend/api => ./api
replace github.com/alexbaily/KitchenHelper-backend/models => ./models

require (
	github.com/alexbaily/KitchenHelper-backend/api v0.0.0-00010101000000-000000000000
	github.com/alexbaily/KitchenHelper-backend/models v0.0.0
	github.com/auth0/go-jwt-middleware v1.0.0 // indirect
	github.com/aws/aws-lambda-go v1.23.0 // indirect
	github.com/aws/aws-sdk-go v1.38.40 // indirect
	github.com/awslabs/aws-lambda-go-api-proxy v0.10.0 // indirect
	github.com/form3tech-oss/jwt-go v3.2.3+incompatible // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
)
