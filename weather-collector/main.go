package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	awslambda "github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"

	"github.com/veselj/dsrc-weather/weather-collector/weather/record"
	"github.com/veselj/dsrc-weather/weather-collector/weather/station"
	"github.com/veselj/dsrc-weather/weather-collector/weather/store"
)

type Event struct {
	Count int `json:"count"`
}

func handler(ctx context.Context, event Event) error {

	weather, err := station.GetWeather()
	if err != nil {
		log.Println("Error getting weather", err)
		return err
	}
	sample := record.AsSample(weather)
	log.Printf("Weather sample: %+v\n", sample)

	err = store.Save(sample)
	if err != nil {
		log.Println("Error saving weather sample", err)
		return err
	}

	if event.Count < 6 {
		//time.Sleep(10 * time.Second)
		//log.Printf("Self-invoking lambda, count: %d\n", event.Count)
		//err = selfInvoke(ctx, event.Count)
	} else {
		log.Println("Reached max self-invoke count.")
	}

	return nil
}

func main() {
	lambda.Start(handler)
}

func selfInvoke(ctx context.Context, count int) error {

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}
	payloadStruct := Event{Count: count + 1}
	payloadBytes, err := json.Marshal(payloadStruct)
	if err != nil {
		return err
	}

	client := awslambda.NewFromConfig(cfg)
	_, err = client.Invoke(ctx, &awslambda.InvokeInput{
		FunctionName:   aws.String(os.Getenv("AWS_LAMBDA_FUNCTION_NAME")),
		InvocationType: types.InvocationTypeEvent,
		Payload:        payloadBytes,
	})
	if err != nil {
		log.Println("Error invoking lambda function", err)
		return err
	}
	return nil
}
