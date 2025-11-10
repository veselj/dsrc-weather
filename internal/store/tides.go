package store

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/veselj/dsrc-weather/weather-collector/weather/tides"
)

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

func (c *DynamoClient) SaveTides(tides []tides.Tide) error {
	err := c.PutTides(context.Background(), tides)
	if err != nil {
		log.Printf("failed to put sample, %+v", err)
		return err
	}
	log.Printf("Saved tides: %+v", tides)
	return nil
}
