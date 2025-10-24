package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/veselj/dsrc-weather/weather-collector/weather/record"
	"github.com/veselj/dsrc-weather/weather-collector/weather/station"
	"github.com/veselj/dsrc-weather/weather-collector/weather/store"
)

func handler(ctx context.Context) error {

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

	return nil
}

func main() {
	lambda.Start(handler)
}
