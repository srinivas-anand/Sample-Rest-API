package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sample-Rest-API/dao"
	"github.com/Sample-Rest-API/utils"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func addOne(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var newMobileItem utils.Mobiles

	//Converting Request body to go value type
	json.Unmarshal(requestBody, &newMobileItem)

	//Converting go value Type to DynamoDb attribute value
	av, err := dynamodbattribute.MarshalMap(newMobileItem)
	if err != nil {
		fmt.Println("Error to convert to DDB AV type:", err.Error())
	}

	tableName := utils.GetConfig().Table
	dao.DDBAddOne(av, tableName)
	json.NewEncoder(w).Encode(newMobileItem)

}
