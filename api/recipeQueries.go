package api

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/alexbaily/KitchenHelper-backend/models"

	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func (d DynamoInt) queryRecipes(UserID string, recipe string, table string) (queryJson []byte) {

	//keyCondition and Projection are required for the expression builder.
	userIDCondition := expression.Key("UserID").Equal(expression.Value(UserID))
	projection := expression.ProjectionBuilder{}

	value := reflect.Indirect(reflect.ValueOf(&models.RecipeRecord{}))
	//Get all of the fields of the models record and add these to the projection.
	for i := 0; i < value.Type().NumField(); i++ {
		projection = projection.AddNames(expression.Name(value.Type().Field(i).Name))
	}

	if recipe == "" {
		expr, err := expression.NewBuilder().
			WithKeyCondition(userIDCondition).
			WithProjection(projection).
			Build()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		recipeCondition := expression.Key("recipeIdentifier").BeginsWith(recipe)
		expr, err := expression.NewBuilder().
			WithKeyCondition(userIDCondition.And(recipeCondition)).
			WithProjection(projection).
			Build()
		if err != nil {
			fmt.Println(err)
		}
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
	result, err := DynaDB.Client.Query(params)
	if err != nil {
		fmt.Println(err)
	}
	//Initilise the slice of RecipeRecord
	recs := []models.RecipeRecord{}

	//UnMarshal the DynamoDB results into a RecipeRecord and store in recs
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &recs)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
	}

	//Marshal the records into JSON
	queryJson, err = json.Marshal(recs)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal records, %v", err))
	}
	log.Printf("records %+v", recs)
	return

}

func addRecipe(UserID string, table string) {
	//Generate a new UUID
	recUUID := uuid.New()

	//Create the UpdateItemInput for updating the DynamoDB table.
	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {
				S: aws.String(UserID),
			},
			"recipeIdentifier": {
				S: aws.String(locName + "#" + recUUID.String()),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#PN": aws.String("ProductName"),
			"#Q":  aws.String("Quantity"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pn": {
				S: aws.String(productName),
			},
			":q": {
				N: aws.String(quantity),
			},
		},
		TableName:        aws.String(table),
		UpdateExpression: aws.String("SET #PN = :pn, #Q = :q"),
	}
	result, err := DynaDB.Client.UpdateItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				fmt.Println(dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeTransactionConflictException:
				fmt.Println(dynamodb.ErrCodeTransactionConflictException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result.ConsumedCapacity)
}
