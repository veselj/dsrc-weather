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
	"github.com/veselj/dsrc-weather/weather-collector/weather/tides"
)

const (
	SamplesTable = "WeatherSamples"
	TidesTable   = "TideTimes"
)

type DynamoClient struct {
	client          *dynamodb.Client
	sampleTableName string
	tidesTableName  string
}

func NewDynamoClient() *DynamoClient {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return &DynamoClient{
		client:          dynamodb.NewFromConfig(cfg),
		sampleTableName: SamplesTable,
		tidesTableName:  TidesTable,
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
		TableName: aws.String(c.sampleTableName),
		Item:      item,
	})
	return err
}

func (c *DynamoClient) PutTides(ctx context.Context, tides []tides.Tide) error {
	t, err := c.GetTides(ctx)
	if err == nil && len(t) > 0 {
		log.Printf("Tides already exist in the database, skipping insert")
		return nil
	}

	for _, tide := range tides {
		item, err := attributevalue.MarshalMap(struct {
			When      int64   `dynamodbav:"When"`
			Type      int     `dynamodbav:"Type"`
			Height    float64 `dynamodbav:"Height"`
			ExpiresAt int64   `dynamodbav:"expires_at"`
		}{
			When:      tide.Time.Unix(),
			Type:      tide.Type,
			Height:    tide.Height,
			ExpiresAt: time.Now().Add(time.Hour * 24 * 14).Unix(), // 14 days expiration
		})
		if err != nil {
			return err
		}

		_, err = c.client.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: aws.String(c.tidesTableName),
			Item:      item,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *DynamoClient) GetTides(ctx context.Context) ([]tides.Tide, error) {
	loc, err := time.LoadLocation("Europe/London")
	if err != nil {
		loc = time.UTC
	}
	now := time.Now().In(loc)
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endOfDay := startOfDay.Add(24 * time.Hour)

	keyCond := "#When BETWEEN :start AND :end"
	exprAttrNames := map[string]string{
		"#When": "When",
	}
	exprAttrValues := map[string]types.AttributeValue{
		":start": &types.AttributeValueMemberN{Value: strconv.FormatInt(startOfDay.Unix(), 10)},
		":end":   &types.AttributeValueMemberN{Value: strconv.FormatInt(endOfDay.Unix()-1, 10)},
	}

	result, err := c.client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String(c.tidesTableName),
		KeyConditionExpression:    aws.String(keyCond),
		ExpressionAttributeNames:  exprAttrNames,
		ExpressionAttributeValues: exprAttrValues,
	})
	if err != nil {
		return nil, err
	}

	var tidesList []tides.Tide
	err = attributevalue.UnmarshalListOfMaps(result.Items, &tidesList)
	if err != nil {
		log.Printf("UnmarshalListOfMaps error: %v", err)
		return nil, err
	}
	return tidesList, nil
}

func (c *DynamoClient) SaveSample(sample *record.Sample) error {
	err := c.PutSample(context.Background(), sample)
	if err != nil {
		log.Printf("failed to put sample, %+v", err)
		return err
	}
	log.Printf("Saved sample: %+v", sample)
	return nil
}

func (c *DynamoClient) SaveTides(tides []tides.Tide) error {
	err := c.PutTides(context.Background(), tides)
	if err != nil {
		log.Printf("failed to put sample, %+v", err)
		return err
	}
	log.Printf("Saved tides: %+v", tides)
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
		TableName:                 aws.String(c.sampleTableName),
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
