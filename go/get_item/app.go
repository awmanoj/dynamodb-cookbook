package main

import(
	"log"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type record struct {
	_type string
	_data string 
}

func main() {
	if len(os.Args) != 3 {
		log.Println("err:", "Usage: ", os.Args[0], " <tableName> <pkValue>")
		return
	}	

	sess, err := session.NewSession()
	if err != nil {
		log.Println("err:", "failed to create session - ", err)
		return
	}

	svc := dynamodb.New(sess)
	
	params := &dynamodb.GetItemInput{

		Key: map[string]*dynamodb.AttributeValue{
			"uid": {
				S: aws.String(os.Args[2]),
			}, 
		},
		ExpressionAttributeNames: map[string]*string{
			"#Sess": aws.String("session"),
		},
		TableName: aws.String(os.Args[1]),
		ConsistentRead: aws.Bool(false),
		ProjectionExpression: aws.String("uid, #Sess, platform"),
		ReturnConsumedCapacity: aws.String("NONE"),
	}

	resp, err := svc.GetItem(params)
	if err != nil {
		log.Println("err:", "retrieving value ", err)
		return
	}

	for key, value := range resp.Item {
		log.Println("::", key, " = ", *value.S)
	}
}

