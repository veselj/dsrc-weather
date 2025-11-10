package store

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/veselj/dsrc-weather/internal/record"
	"github.com/veselj/dsrc-weather/weather-collector/weather/station"
)

func (c *DynamoClient) PutWeather(
	ctx context.Context,
	stationData *station.WeatherData) error {
	wDetails := record.AsWeatherData(stationData)
	item, err := attributevalue.MarshalMap(struct {
		Bucket            string  `dynamodbav:"Bucket"`
		WindSpeed         float64 `dynamodbav:"WindSpeed"`
		Temperature       float64 `dynamodbav:"Temperature"`
		FeelsLike         float64 `dynamodbav:"FeelsLike"`
		WindDirection     int     `dynamodbav:"WindDirection"`
		WindDirectionName string  `dynamodbav:"WindDirectionName"`
		Barometer         float64 `dynamodbav:"Barometer"`
		BarometerUnits    string  `dynamodbav:"BarometerUnits"`
		BarometerTrend    string  `dynamodbav:"BarometerTrend"`
		Rain              float64 `dynamodbav:"Rain"`
		RainUnits         string  `dynamodbav:"RainUnits"`
		ChanceOfRain      int     `dynamodbav:"ChanceOfRain"`
		Humidity          float64 `dynamodbav:"Humidity"`
		Forecast          string  `dynamodbav:"Forecast"`
		ExpiresAt         int64   `dynamodbav:"expires_at"`
	}{
		Bucket:            wDetails.Bucket,
		WindSpeed:         wDetails.WindSpeed,
		Temperature:       wDetails.Temperature,
		FeelsLike:         wDetails.FeelsLike,
		WindDirection:     wDetails.WindDirection,
		WindDirectionName: wDetails.WindDirectionName,
		Barometer:         wDetails.Barometer,
		BarometerUnits:    wDetails.BarometerUnits,
		BarometerTrend:    wDetails.BarometerTrend,
		Rain:              wDetails.Rain,
		RainUnits:         wDetails.RainUnits,
		ChanceOfRain:      wDetails.ChanceOfRain,
		Humidity:          wDetails.Humidity,
		Forecast:          wDetails.Forecast,
		ExpiresAt:         time.Now().Add(time.Hour * 24 * 14).Unix(), // 14 days expiration
	})
	if err != nil {
		return err
	}

	_, err = c.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(c.weatherTableName),
		Item:      item,
	})
	return err
}

func (c *DynamoClient) GetWeather(ctx context.Context) (*record.WeatherDetails, error) {
	now := time.Now()
	partitionKeyValue := now.Format(record.BucketFormat)

	result, err := c.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(c.weatherTableName),
		Key: map[string]types.AttributeValue{
			"Bucket": &types.AttributeValueMemberS{Value: partitionKeyValue},
		},
	})
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}
	var weather record.WeatherDetails
	err = attributevalue.UnmarshalMap(result.Item, &weather)
	if err != nil {
		return nil, err
	}
	return &weather, nil
}
