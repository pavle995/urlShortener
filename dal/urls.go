package dal

import (
	"errors"
	"urlShortener/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	log "github.com/sirupsen/logrus"
)

type DynamoDBClient struct {
	client    *dynamodb.DynamoDB
	tableName string
}

var dbClient *DynamoDBClient

func GetDbClient() *DynamoDBClient {
	if dbClient == nil {
		db := newDbClient()
		dbClient = &db
	}
	return dbClient
}

func newDbClient() DynamoDBClient {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	return DynamoDBClient{client: svc, tableName: "urls"}
}

func (c *DynamoDBClient) InsertNewRecord(url *models.Url) error {
	av, err := dynamodbattribute.MarshalMap(*url)
	if err != nil {
		log.Error("MarshalMap error: " + err.Error())
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(c.tableName),
	}

	_, err = c.client.PutItem(input)
	if err != nil {
		log.Error("fail writing to db: " + err.Error())
		return err
	}

	return nil
}

func (c *DynamoDBClient) GetRedirect(id int) (*string, error) {
	filt := expression.Name("id").Equal(expression.Value(id))
	proj := expression.NamesList(expression.Name("fullUrl"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		log.Error("error building a scan expression: " + err.Error())
		return nil, err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(c.tableName),
	}

	result, err := c.client.Scan(params)
	if err != nil {
		log.Error("error scanning from db: " + err.Error())
		return nil, err
	}

	for _, i := range result.Items {
		item := models.Url{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			log.Error("error UnmarshalingMap: " + err.Error())
			return nil, err
		}

		return &item.Url, nil
	}

	return nil, errors.New("not found")
}
