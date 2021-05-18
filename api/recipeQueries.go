package api

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/alexbaily/KitchenHelper-backend/models"

	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func (d DynamoInt) queryRecipes(UserID string, recipe string, table string) (queryJson []byte, err error) {

	//keyCondition and Projection are required for the expression builder.
	userIDCondition := expression.Key("UserID").Equal(expression.Value(UserID))
	projection := expression.ProjectionBuilder{}

	value := reflect.Indirect(reflect.ValueOf(&models.RecipeRecord{}))
	//Get all of the fields of the models record and add these to the projection.
	for i := 0; i < value.Type().NumField(); i++ {
		projection = projection.AddNames(expression.Name(value.Type().Field(i).Name))
	}

	var expr expression.Expression
	//If recipe is empty then return all of the recipes for the user.
	if recipe == "" {
		expr, err = expression.NewBuilder().
			WithKeyCondition(userIDCondition).
			WithProjection(projection).
			Build()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		recipeCondition := expression.Key("recipeIdentifier").BeginsWith(recipe)
		expr, err = expression.NewBuilder().
			WithKeyCondition(userIDCondition.And(recipeCondition)).
			WithProjection(projection).
			Build()
		if err != nil {
			fmt.Println(err)
			return nil, err
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
	result, err := d.Client.Query(params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//Initilise the slice of RecipeRecord
	recs := []models.RecipeRecord{}

	//Iterate over the items so that we can unmarshal them
	for _, i := range result.Items {
		rec := models.RecipeRecord{}

		err = dynamodbattribute.UnmarshalMap(i, &rec)
		fmt.Println(i)

		if err != nil {
            panic(fmt.Sprintf("Got error unmarshalling: %v", err))
			return nil, err
        }

		recs = append(recs, rec)
	}

	//UnMarshal the DynamoDB results into a RecipeRecord and store in recs
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &recs)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
		return nil, err
	}

	//Marshal the records into JSON
	queryJson, err = json.Marshal(recs)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal records, %v", err))
		return nil, err
	}
	return queryJson, err

}

func (d DynamoInt) addRecipe(UserID string, recipe models.RecipeRecord, table string) error {
	
	//Generate a new UUID
	recipeUUID := uuid.New()

	//Convert the steps to a list of maps
	lMap, err := dynamodbattribute.MarshalList(recipe.Steps)
	if err != nil {
		panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
	}
	//Convert the ingredients to a map[string]*dynamodb.AttributeValue
	ingredients, err := dynamodbattribute.MarshalList(recipe.Ingredients)
	if err != nil {
		panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
	}
	//Create the UpdateItemInput for updating the DynamoDB table.
	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {
				S: aws.String(UserID),
			},
			"recipeIdentifier": {
				S: aws.String(recipeUUID.String()),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#RN": aws.String("RecipeName"),
			"#D":  aws.String("Description"),
			"#P": aws.String("PhotoURL"),
			"#N":  aws.String("Notes"),
			"#S":  aws.String("Sharing"),
			"#ST": aws.String("Steps"),
			"#I": aws.String("Ingredients"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":rn": {
				S: aws.String(recipe.RecipeName),
			},
			":d": {
				S: aws.String(recipe.Description),
			},
			":p": {
				S: aws.String(recipe.PhotoURL),
			},
			":n": {
				S: aws.String(recipe.Notes),
			},
			":s": {
				S: aws.String(recipe.Sharing),
			},
			":st": {
				L: lMap,
			},
			":i": {
				L: ingredients,
			},
		},
		TableName:        aws.String(table),
		UpdateExpression: aws.String(
			"SET #RN = :rn, #D = :d, #P = :p, #N = :n, #S = :s, #ST = :st, #I = :i"),
	}
	_, err = d.Client.UpdateItem(input)
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
		return err
	}
	return nil
}
