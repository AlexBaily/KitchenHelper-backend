# KitchenHelper-backend

This is the backend API for the KitchenHelper app. 

Written in Go lang utilizing mux and negroni for http calls, Auth0 and JWT for authentication and DynamoDB for NoSQL.

Environment variables required:

*	`KITCHTABLE` - The DynamoDB table that is to be used for locations.
*   `RECIPETABLE` - The DynamoDB table that is used for recipes.
*	`TOKENURL` - The Auth0 Token URL used to verify JWT tokens.
*	`AUTH_AUDIENCE` - The Auth0 apis audience, used to verify JWT tokens.
*   `AWS_DEFAULT_REGION` - The AWS region in which the DynamoDB tables are located.


## Build

`go get -d -v ./...`
`go install -v ./...`

## Docker Build and run

`docker build . -t $image-tag`
`docker run -d -p 8080:8080 -e KITCHTABLE=$var -e AUTH_AUDIENCE=$var -e TOKENURL=$var -e AWS_DEFAULT_REGION=$var -e RECIPETABLE=$var $image-tag` 