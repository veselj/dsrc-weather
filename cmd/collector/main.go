package main

import (
	"github.com/veselj/dsrc-weather/internal/record"
	"github.com/veselj/dsrc-weather/weather-collector/weather/station"
	"log"
	"os"
)

func main() {
	weather, err := station.GetWeather()
	if err != nil {
		log.Println("Error getting weather", err)
		os.Exit(-1)
	}
	sample := record.AsSample(weather)
	log.Printf("Weather sample: %+v\n", sample)
}
