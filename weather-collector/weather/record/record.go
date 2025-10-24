package record

import (
	"strconv"
	"time"

	"github.com/veselj/dsrc-weather/weather-collector/weather/station"
)

func AsSample(w *station.WeatherData) *Sample {
	temper := asDegrees(w.Temperature)
	feelsLike := asDegrees(w.TemperatureFeelLike)
	now := time.Now()
	rv := &Sample{
		Wind:        asKnots(w.Wind, w.WindUnits),
		Direction:   w.WindDirection,
		Temperature: temper,
		FeelsLike:   feelsLike,
		Bucket:      asBucket(now),
		When:        now.Unix(),
	}
	return rv
}

// Hourly buckets
func asBucket(t time.Time) string {
	return t.Format("2006010215")
}

func asDegrees(temper string) float64 {
	value, err := strconv.ParseFloat(temper, 64)
	if err != nil {
		return 0
	}
	return value
}

func asKnots(mph string, windUnit string) float64 {
	if windUnit == "mph" {
		speed, err := strconv.ParseFloat(mph, 64)
		if err != nil {
			return 0
		}
		speedKn := speed * 0.8689758
		return speedKn
	}
	return 0
}
