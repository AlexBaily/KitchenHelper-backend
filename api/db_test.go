package api

import (
	"bytes"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type mockedDynamo struct {
	dynamodbiface.DynamoDBAPI
	Resp    *dynamodb.QueryOutput
	addResp *dynamodb.UpdateItemOutput
}

func (m mockedDynamo) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return m.Resp, nil
}

func (m mockedDynamo) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return m.addResp, nil
}

func TestQueryLocations(t *testing.T) {
	cases := []struct {
		Resp     *dynamodb.QueryOutput
		Expected []byte
	}{
		{
			Resp: &dynamodb.QueryOutput{
				Items: []map[string]*dynamodb.AttributeValue{
					{
						"ProductIdentifier": &dynamodb.AttributeValue{
							S: aws.String("kitchen#tomatoes"),
						},
						"Location": &dynamodb.AttributeValue{
							S: aws.String("kitchen"),
						},
					},
				},
			},
			Expected: []byte("{\"locations\":[\"kitchen\"]}"),
		},
	}

	for i, c := range cases {
		db := DynamoInt{
			Client: mockedDynamo{Resp: c.Resp},
		}
		rows := db.queryLocations("testuserid", "testtable")
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

func TestAddLocation(t *testing.T) {
	cases := []struct {
		addResp  *dynamodb.UpdateItemOutput
		Expected int
	}{
		{
			addResp: &dynamodb.UpdateItemOutput{
				Attributes: map[string]*dynamodb.AttributeValue{
					"Location": &dynamodb.AttributeValue{
						S: aws.String("fridge"),
					},
				},
			},
			Expected: 200,
		},
	}

	for i, c := range cases {
		db := DynamoInt{
			Client: mockedDynamo{addResp: c.addResp},
		}
		rows := db.addLocation("10000000", "pantry-table", "fridge")
		if a, e := rows, c.Expected; a != e {
			t.Fatalf("%d, expected status of %d, got %d", i, e, a)
		}
	}

}

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
			Expected: []byte("{\"recipes\":[\"bean burritos\"]}"),
		},
	}

	for i, c := range cases {
		db := DynamoInt{
			Client: mockedDynamo{Resp: c.Resp},
		}
		rows := db.queryRecipes("testuserid", "recipestable")
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
