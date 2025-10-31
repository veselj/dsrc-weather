package store

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/veselj/dsrc-weather/internal/record"
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
		Bt        string  `dynamodbav:"Bt"`
		Wd        float64 `dynamodbav:"Wd"`
		Dn        int     `dynamodbav:"Dn"`
		Te        float64 `dynamodbav:"Te"`
		Fl        float64 `dynamodbav:"Fl"`
		Wn        int64   `dynamodbav:"Wn"`
		ExpiresAt int64   `dynamodbav:"expires_at"`
	}{
		Bt:        s.Bt,
		Wd:        s.Wd,
		Dn:        s.Dn,
		Te:        s.Te,
		Fl:        s.Fl,
		Wn:        s.Wn,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 14).Unix(), // 14 days expiration
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
	log.Printf("Saved sample: %+v", sample)
	return nil
}

func (c *DynamoClient) Samples(ctx context.Context, fromUnix int64) []record.Sample {
	// Query parameters
	partitionKeyValue := time.Unix(fromUnix, 0).Format(record.BucketFormat)
	startRangeKey := fromUnix

	// Build expression for PK and range key between two values
	keyCond := "#PK = :pk AND #SK >= :startSK"
	exprAttrNames := map[string]string{
		"#PK": "Bt",
		"#SK": "Wn",
	}
	exprAttrValues := map[string]types.AttributeValue{
		":pk":      &types.AttributeValueMemberS{Value: partitionKeyValue},
		":startSK": &types.AttributeValueMemberN{Value: strconv.FormatInt(startRangeKey, 10)},
	}

	// Execute query
	result, err := c.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String(c.tableName),
		KeyConditionExpression:    aws.String(keyCond),
		ExpressionAttributeNames:  exprAttrNames,
		ExpressionAttributeValues: exprAttrValues,
	})
	if err != nil {
		log.Fatalf("Query API call failed: %v", err)

	}
	samples := []record.Sample{}
	err = attributevalue.UnmarshalListOfMaps(result.Items, &samples)
	if err != nil {
		log.Printf("failed to unmarshal items: %v", err)
	}
	return samples
}

func GetHourlyBucketStarts(fromUnix int64) []int64 {
	startBucketTime := time.Unix(fromUnix, 0).Truncate(time.Hour)
	fromUnixTime := time.Unix(fromUnix, 0)
	now := time.Now().Truncate(time.Hour)
	duration := now.Sub(startBucketTime)
	hours := int(duration.Hours())
	fromTimes := make([]int64, 0, hours+1)
	fromTimes = append(fromTimes, fromUnix)
	for i := 1; i <= hours; i++ {
		t := fromUnixTime.Add(time.Duration(i) * time.Hour).Truncate(time.Hour).Unix()
		fromTimes = append(fromTimes, t)
	}
	return fromTimes
}
