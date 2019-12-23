package api

import (
	"fmt"
        "log"
        "strings"
	"encoding/json"
	"KitchenHelper-backend/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func queryLocations(UserID string, table string)(queryJson []byte) {
        sess := session.Must(session.NewSessionWithOptions(session.Options{
                SharedConfigState: session.SharedConfigEnable,
        }))

        svc := dynamodb.New(sess)



        //keyCondition and Projection are required for the expression builder.
        keyCondition  := expression.Key("UserID").Equal(expression.Value(UserID))
        projection := expression.NamesList(
                expression.Name("productIdentifier"),
        )
        expr, err := expression.NewBuilder().
                WithKeyCondition(keyCondition).
                WithProjection(projection).
                Build()
        if err != nil {
                fmt.Println(err)
        }
        //Load up the parameters into a struct
        params := &dynamodb.QueryInput{
                KeyConditionExpression:    expr.KeyCondition(),
                ExpressionAttributeNames:  expr.Names(),
                ExpressionAttributeValues: expr.Values(),
                ProjectionExpression:      expr.Projection(),
                TableName:                 aws.String(table),
        }

        //Complete a query of the table with the params from above
        result, err := svc.Query(params)
        if err != nil {
                fmt.Println(err)
        }
        //Initilise the slice of LocRecord
        recs := []models.LocRecord{}

        //UnMarshal the DynamoDB results into a LocRecord and store in recs 
        err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &recs)
        if err != nil {
                panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
        }
        
        //Put the locations into a map so that we can grab each unique location.
        //Use a map so it's O(1) and we don't have to iterate through the entire array
        //each time that we want to check on insert.
        locationsMap := make(map[string]struct{})

        for i := range(recs) {
                locationEnd := strings.Index(recs[i].ProductIdentifier, "#")
                fmt.Printf("%v %v", recs[i], locationEnd)
                locationsMap[recs[i].ProductIdentifier[0:locationEnd]] = struct{}{}
        }

        //Convert the map to a slice so we can convert to a json object.
        var locationsSlice []string
        for k, _ := range locationsMap {
                locationsSlice = append(locationsSlice, k)
        }
        //Convert into the final map to convert to JSON for writing back to the client.
        locationsFinalMap := make(map[string][]string)
        locationsFinalMap["locations"] = locationsSlice

        //Marshal the records into JSON
        queryJson, err = json.Marshal(locationsFinalMap)
        if err != nil {
                panic(fmt.Sprintf("failed to marshal records, %v", err))
        }
        log.Printf("records %+v", recs[0])
        return

}


/*
func queryProducts(UserID string, location string,table string)(queryJson []byte) {
        sess := session.Must(session.NewSessionWithOptions(session.Options{
                SharedConfigState: session.SharedConfigEnable,
        }))

        svc := dynamodb.New(sess)

        //Filter and Projection are required for the expression builder.
        filter := expression.Name("UserID").Equal(expression.Value(UserID))
        projection := expression.NamesList(
                expression.Name("productIdentifier"),
                expression.Name("ProductName"),
                expression.Name("Quantity"),
        )
        expr, err := expression.NewBuilder().
                WithFilter(filter).
                WithProjection(projection).
                Build()
        if err != nil {
                fmt.Println(err)
        }
        //Load up the parameters into a struct
        params := &dynamodb.ScanInput{
                ExpressionAttributeNames:  expr.Names(),
                ExpressionAttributeValues: expr.Values(),
                FilterExpression:          expr.Filter(),
                ProjectionExpression:      expr.Projection(),
                TableName:                 aws.String(table),
        }

        //Complete a scan of the table with the params from above
        result, err := svc.Scan(params)
        if err != nil {
                fmt.Println(err)
        }

        recs := []models.LocRecord{}

        err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &recs)
        if err != nil {
                panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
        }

        //Marshal the records into JSON
        queryJson, err = json.Marshal(recs[0])
        if err != nil {
                panic(fmt.Sprintf("failed to marshal records, %v", err))
        }
        log.Printf("records %+v", recs[0])
        return

}*/
