package store

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/veselj/dsrc-weather/weather-collector/weather/record"
)

const (
	SamplesTable = "WeatherSamples"
)

type DynamoClient struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoClient() *DynamoClient {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return &DynamoClient{
		client:    dynamodb.NewFromConfig(cfg),
		tableName: SamplesTable,
	}
}

func (c *DynamoClient) PutSample(ctx context.Context, s *record.Sample) error {
	item, err := attributevalue.MarshalMap(struct {
		Bucket      string  `dynamodbav:"Bucket"`
		Wind        float64 `dynamodbav:"Wind"`
		Direction   int     `dynamodbav:"Direction"`
		Temperature float64 `dynamodbav:"Temperature"`
		FeelsLike   float64 `dynamodbav:"FeelsLike"`
		When        int64   `dynamodbav:"When"`
	}{
		Bucket:      s.Bucket,
		Wind:        s.Wind,
		Direction:   s.Direction,
		Temperature: s.Temperature,
		FeelsLike:   s.FeelsLike,
		When:        s.When,
	})
	if err != nil {
		return err
	}

	_, err = c.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(c.tableName),
		Item:      item,
	})
	return err
}

func Save(sample *record.Sample) error {
	dynamo := NewDynamoClient()
	err := dynamo.PutSample(context.Background(), sample)
	if err != nil {
		log.Printf("failed to put sample, %+v", err)
		return err
	}
	log.Printf("Savedsample: %+v", sample)
	return nil
}
