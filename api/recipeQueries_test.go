package api

import (
	"bytes"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestQueryRecipes(t *testing.T) {
	cases := []struct {
		Resp     *dynamodb.QueryOutput
		Expected []byte
	}{
		{
			Resp: &dynamodb.QueryOutput{
				Items: []map[string]*dynamodb.AttributeValue{
					{
						"RecipeIdentifier": &dynamodb.AttributeValue{
							S: aws.String("bean burritos"),
						},
					},
				},
			},
			Expected: []byte("[{\"userID\":\"\",\"recipeIdentifier\":\"bean burritos\",\"recipeName\":\"\",\"photoURL\":\"\",\"notes\":\"\",\"description\":\"\",\"sharing\":\"\",\"steps\":null,\"ingredients\":null}]"),
		},
	}

	for i, c := range cases {
		db := DynamoInt{
			Client: mockedDynamo{Resp: c.Resp},
		}
		rows, _ := db.queryRecipes("testuserid", "","recipestable")
		if a, e := rows, c.Expected; bytes.Compare(a, e) != 0 {
			t.Fatalf("%d, expected %d row, got %d", i, e, a)
		}
		for j, rows := range rows {
			if a, e := rows, c.Expected[j]; a != e {
				t.Errorf("%d, expected %v row, got %v", i, e, a)
			}
		}
	}

}

func TestQueryRecipe(t *testing.T) {
	cases := []struct {
		Resp     *dynamodb.QueryOutput
		Expected []byte
	}{
		{
			Resp: &dynamodb.QueryOutput{
				Items: []map[string]*dynamodb.AttributeValue{
					{
						"RecipeIdentifier": &dynamodb.AttributeValue{
							S: aws.String("bean burritos"),
						},
						"PhotoURL": &dynamodb.AttributeValue{
							S: aws.String("http://here-is-a-photo-s3.aws.photo"),
						},
						"Description": &dynamodb.AttributeValue{
							S: aws.String("Test"),
						},
					},
				},
			},
			Expected: []byte("[{\"userID\":\"\",\"recipeIdentifier\":\"bean burritos\",\"recipeName\":\"\",\"photoURL\":\"http://here-is-a-photo-s3.aws.photo\",\"notes\":\"\",\"description\":\"Test\",\"sharing\":\"\",\"steps\":null,\"ingredients\":null}]"),
		},
	}

	for i, c := range cases {
		db := DynamoInt{
			Client: mockedDynamo{Resp: c.Resp},
		}
		rows, _ := db.queryRecipes("testuserid", "bean", "recipestable")
		if a, e := rows, c.Expected; bytes.Compare(a, e) != 0 {
			t.Fatalf("%d, expected %d row, got %d", i, e, a)
		}
		for j, rows := range rows {
			if a, e := rows, c.Expected[j]; a != e {
				t.Errorf("%d, expected %v row, got %v", i, e, a)
			}
		}
	}

}
