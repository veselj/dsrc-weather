package store

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/veselj/dsrc-weather/weather-collector/weather/tides"
)

const (
	TideBucketFormat = "20060102"
)

// Daily buckets
func AsTidesBucket(t int64) string {
	return time.Unix(t, 0).Format(TideBucketFormat)
}

func (c *DynamoClient) PutTides(ctx context.Context, tides []tides.Tide) error {
	t, err := c.GetTides(ctx)
	if err == nil && len(t) > 0 {
		log.Printf("Tides already exist in the database, skipping insert")
		return nil
	}

	for _, tide := range tides {
		item, err := attributevalue.MarshalMap(struct {
			Bucket    string  `dynamodbav:"Bucket"`
			Time      int64   `dynamodbav:"Time"`
			Type      int     `dynamodbav:"Type"`
			Height    float64 `dynamodbav:"Height"`
			ExpiresAt int64   `dynamodbav:"expires_at"`
		}{
			Bucket:    AsTidesBucket(tide.Time),
			Time:      tide.Time,
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

	keyCond := "#Bucket = :bucket"
	exprAttrNames := map[string]string{
		"#Bucket": "Bucket",
	}
	exprAttrValues := map[string]types.AttributeValue{
		":bucket": &types.AttributeValueMemberS{Value: AsTidesBucket(now.Unix())},
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

func (c *DynamoClient) SaveTides(tides []tides.Tide) error {
	err := c.PutTides(context.Background(), tides)
	if err != nil {
		log.Printf("failed to put sample, %+v", err)
		return err
	}
	log.Printf("Saved tides: %+v", tides)
	return nil
}
