# KitchenHelper-backend

This is the backend API for the KitchenHelper app. 

Written in Go lang utilizing mux and negroni for http calls, Auth0 and JWT for authentication and DynamoDB for NoSQL.

Environment variables required:

*	`KITCHTABLE` - The DynamoDB table that is to be used for the recipes/locations.
*	`TOKENURL` - The Auth0 Token URL used to verify JWT tokens.
*	`AUTH_AUDIENCE` - The Auth0 apis audience, used to verify JWT tokens.
