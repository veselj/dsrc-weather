package record

type Sample struct {
	Wd float64
	Dn int
	Te float64
	Fl float64
	Wn int64 // Unix timestamp
	Bt string
}

const (
	BucketFormat = "2006010215"
)

type WeatherDetails struct {
	Bucket            string
	WindSpeed         float64
	Temperature       float64
	FeelsLike         float64
	WindDirection     int
	WindDirectionName string
	Barometer         float64
	BarometerUnits    string
	BarometerTrend    string
	Rain              float64
	RainUnits         string
	ChanceOfRain      int
	Humidity          float64
	Forecast          string
}
