package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/veselj/dsrc-weather/internal/record"
	"github.com/veselj/dsrc-weather/internal/store"
	"github.com/veselj/dsrc-weather/weather-collector/weather/tides"

	"github.com/veselj/dsrc-weather/weather-collector/weather/station"
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

	dynClient := store.NewDynamoClient()
	err = dynClient.SaveSample(sample)
	if err != nil {
		log.Println("Error saving weather sample", err)
		return err
	}

	tides, err := tides.Scrape()
	if err != nil {
		log.Println("Error scraping tides", err)
		return err
	}
	err = dynClient.SaveTides(tides)
	if err != nil {
		log.Println("Error saving tides", err)
		return err
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
