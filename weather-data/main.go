package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/veselj/dsrc-weather/internal/store"
	"github.com/veselj/dsrc-weather/weather-collector/weather/record"
)

func handler(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {

	from := time.Now().Add(-time.Hour * 3)
	samples := retrieveSamples(ctx, from.Unix())
	jsonBody, err := json.Marshal(samples)
	if err != nil {
		jsonBody = []byte(err.Error())
	}
	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: string(jsonBody),
	}, nil
}

func main() {
	lambda.Start(handler)
}

func retrieveSamples(ctx context.Context, fromUnix int64) []record.Sample {
	dynamo := store.NewDynamoClient()
	startTimes := store.GetHourlyBucketStarts(fromUnix)
	var retrieved []record.Sample
	for _, startTime := range startTimes {
		s := dynamo.Samples(ctx, startTime)
		retrieved = append(retrieved, s...)
	}
	return retrieved
}
