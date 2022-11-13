package dal

import (
	"urlShortener/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBClient struct {
	client    *dynamodb.DynamoDB
	tableName string
}

func (c *DynamoDBClient) New(tableName string) DynamoDBClient {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	return DynamoDBClient{client: svc}
}

func (c *DynamoDBClient) InsertNewRecord(url *models.Url) error {
	av, err := dynamodbattribute.MarshalMap(*url)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(c.tableName),
	}

	_, err = c.client.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
