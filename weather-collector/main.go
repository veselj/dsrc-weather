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

	weatherDetail := record.AsWeatherData(weather)
	log.Printf("Weather details: %+v\n", weatherDetail)
	err = dynClient.PutWeather(ctx, weather)
	if err != nil {
		log.Println("Error saving weather details", err)
		return err
	}

	tideTimes, err := tides.Scrape()
	if err != nil {
		log.Println("Error scraping tides", err)
		return err
	}
	log.Printf("Tides: %+v\n", tideTimes)
	err = dynClient.SaveTides(tideTimes)
	if err != nil {
		log.Println("Error saving tides", err)
		return err
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
