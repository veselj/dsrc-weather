package store

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const (
	SamplesTable = "WeatherSamples"
	TidesTable   = "TideTimes"
	WeatherTable = "Weather"
)

type DynamoClient struct {
	client           *dynamodb.Client
	sampleTableName  string
	tidesTableName   string
	weatherTableName string
}

func NewDynamoClient() *DynamoClient {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return &DynamoClient{
		client:           dynamodb.NewFromConfig(cfg),
		sampleTableName:  SamplesTable,
		tidesTableName:   TidesTable,
		weatherTableName: WeatherTable,
	}
}
