# Simple Rest Api using Golang Gorilla mux Router

This Api has following methods to query,scan and put item into a DynamoDB Table

1) /getone/{mid}  (GET)

2) /getall        (GET)

3) /addone        (POST)

            {"Mid":"","Make":"","Model":"","year":0}

## Steps to build

### Pre-req

1) DynamoDB Table

        {"Mid":"","Make":"","Model":"","year":0}

2) IAM role to call DynamoDb

### Run

1)Clone this Repo

2)Export the AWS access keys

3)Run

     go run main.go

4)Run curl commands to Play

        ➜curl http://localhost:9898/getone/1
            {"Mid":"","Make":"","Model":"","year":0}

        ➜ curl http://localhost:9898/getone/1                            
            {"Mid":"1","Make":"Apple","Model":"11 Max Pro","year":2020}  

        ➜ curl http://localhost:9898/getall         
            [{"Mid":"1","Make":"Apple","Model":"11 Max Pro","year":2020}]
        ➜ curl -X POST http://localhost:9898/addone -d '{"Mid":"2","Make":"Apple",         "Model":"X","year":2017}'
            {"Mid":"2","Make":"Apple","Model":"X","year":2017}
        ➜ curl http://localhost:9898/getall                                                                
            [{"Mid":"2","Make":"Apple","Model":"X","year":2017},{"Mid":"1","Make":"Apple","Model":"11 Max Pro","year":2020}]
