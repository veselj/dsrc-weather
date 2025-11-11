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
		Wd: asKnots(w.Wind, w.WindUnits),
		Dn: w.WindDirection,
		Te: temper,
		Fl: feelsLike,
		Bt: asBucket(now),
		Wn: now.Unix(),
	}
	return rv
}

// Hourly buckets
func asBucket(t time.Time) string {
	return t.Format(BucketFormat)
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

func asWindDirection(degrees int) string {
	if degrees < 0 || degrees > 360 {
		return ""
	}
	// multiply by 4 to stay in integer realm (as if 0..1440 degrees) and 45 is then 4*11.25 ...
	segment := ((degrees*4 + 45) % 1440) / 90 // 16 segments
	segmentNames := []string{
		"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW",
	}
	return segmentNames[segment]
}

func asRain(rain string) float64 {
	value, err := strconv.ParseFloat(rain, 64)
	if err != nil {
		return 0
	}
	return value
}

// returns weather description and chance of rain
func partOfDayForecast(w *station.WeatherData) (string, int) {
	if len(w.ForecastOverview) < 1 {
		return "", 0
	}
	london, _ := time.LoadLocation("Europe/London")
	now := time.Now().In(london)
	hour := now.Hour()
	if hour < 6 || hour > 22 {
		return w.ForecastOverview[0].Night.WeatherDesc, w.ForecastOverview[0].Night.Chanceofrain
	} else if hour < 12 {
		return w.ForecastOverview[0].Morning.WeatherDesc, w.ForecastOverview[0].Morning.Chanceofrain
	} else if hour < 17 {
		return w.ForecastOverview[0].Afternoon.WeatherDesc, w.ForecastOverview[0].Afternoon.Chanceofrain
	} else {
		return w.ForecastOverview[0].Evening.WeatherDesc, w.ForecastOverview[0].Evening.Chanceofrain
	}
}

func AsWeatherData(w *station.WeatherData) *WeatherDetails {
	barometer, _ := strconv.ParseFloat(w.Barometer, 64)
	humidity, _ := strconv.ParseFloat(w.Humidity, 64)
	forecast, rainForecast := partOfDayForecast(w)

	return &WeatherDetails{
		Bucket:            asBucket(time.Now()),
		When:              w.LastReceived,
		WindSpeed:         asKnots(w.Wind, w.WindUnits),
		Temperature:       asDegrees(w.Temperature),
		FeelsLike:         asDegrees(w.TemperatureFeelLike),
		WindDirection:     w.WindDirection,
		WindDirectionName: asWindDirection(w.WindDirection),
		Barometer:         barometer,
		BarometerUnits:    w.BarometerUnits,
		BarometerTrend:    w.BarometerTrend,
		Rain:              asRain(w.Rain),
		RainUnits:         w.RainUnits,
		ChanceOfRain:      rainForecast,
		Humidity:          humidity,
		Forecast:          forecast,
	}
}
