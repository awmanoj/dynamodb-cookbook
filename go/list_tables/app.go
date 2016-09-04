package main

import(
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		log.Println("err:", "failed to create session - ", err)
		return
	}

	svc := dynamodb.New(sess)
	
	log.Println("info:", "ddb service reference: ", svc)
	
	params := &dynamodb.ListTablesInput {
		ExclusiveStartTableName: aws.String("TableName"),
		Limit: aws.Int64(100),
	}

	resp, err := svc.ListTables(params)
	if err != nil {
		log.Println("err:", "error retrieving list of tables ", err)
		return
	}

	log.Println(resp)
}

