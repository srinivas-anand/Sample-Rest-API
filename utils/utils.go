package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Configuration struct {
	Table  string
	Region string
}

type Mobiles struct {
	Mid   string `json:"Mid"`
	Make  string `json:"Make"`
	Model string `json:"Model"`
	Year  int    `json:"year"`
}

var Svc *dynamodb.DynamoDB

func GetSession() *session.Session {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(GetConfig().Region)}))
	return sess
}

func DDBClient() *dynamodb.DynamoDB {
	if Svc == nil {
		Svc = dynamodb.New(GetSession())
	}
	return Svc
}

//GetConfig reads the json file ,unmarshalls to the go type
func GetConfig() Configuration {

	pwd, _ := os.Getwd()
	cFile, err := ioutil.ReadFile(pwd + "/configuration/properties.json")
	if err != nil {
		fmt.Println("some error in reading properties json file")
		panic(err.Error())
	}

	var configuration Configuration
	err = json.Unmarshal(cFile, &configuration)
	if err != nil {
		fmt.Println("Unable to convert to Go type Configuration")
		panic(err.Error())
	}

	return configuration
}
