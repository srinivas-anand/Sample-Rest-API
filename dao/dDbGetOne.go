package dao

import (
	"fmt"

	"github.com/Sample-Rest-API/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func DDBGetOne(mid string, targetTable string) utils.Mobiles {

	svc := utils.DDBClient()
	item := utils.Mobiles{}
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: &targetTable,
		Key: map[string]*dynamodb.AttributeValue{
			"Mid": {
				S: aws.String(mid),
			},
		},
	})

	if err != nil {
		fmt.Printf("Failed to connect, %v", err.Error())
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		fmt.Printf("Failed to Unmarshall Record, %v", err.Error())
	}

	if item.Mid == "" {
		fmt.Printf("No items found for the Mid: %s in the table: %s", mid, targetTable)
	}

	return item

}
