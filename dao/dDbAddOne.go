package dao

import (
	"fmt"

	"github.com/Sample-Rest-API/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DDBAddOne(av map[string]*dynamodb.AttributeValue, targetTable string) {

	svc := utils.DDBClient()
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(targetTable),
	}

	_, err := svc.PutItem(input)
	if err != nil {
		fmt.Printf("##Got an error calling PutItem %s ##", av)
		panic(err.Error())
	}
}
