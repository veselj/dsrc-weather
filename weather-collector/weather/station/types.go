package station

type WeatherData struct {
	WindDirection    int `json:"windDirection"`
	ForecastOverview []struct {
		Date    string `json:"date"`
		Morning struct {
			WeatherCode    int     `json:"weatherCode"`
			WeatherDesc    string  `json:"weatherDesc"`
			WeatherIconURL string  `json:"weatherIconUrl"`
			Temp           int     `json:"temp"`
			Chanceofrain   int     `json:"chanceofrain"`
			RainInInches   float64 `json:"rainInInches"`
		} `json:"morning"`
		Afternoon struct {
			WeatherCode    int     `json:"weatherCode"`
			WeatherDesc    string  `json:"weatherDesc"`
			WeatherIconURL string  `json:"weatherIconUrl"`
			Temp           int     `json:"temp"`
			Chanceofrain   int     `json:"chanceofrain"`
			RainInInches   float64 `json:"rainInInches"`
		} `json:"afternoon"`
		Evening struct {
			WeatherCode    int     `json:"weatherCode"`
			WeatherDesc    string  `json:"weatherDesc"`
			WeatherIconURL string  `json:"weatherIconUrl"`
			Temp           int     `json:"temp"`
			Chanceofrain   int     `json:"chanceofrain"`
			RainInInches   float64 `json:"rainInInches"`
		} `json:"evening"`
		Night struct {
			WeatherCode    int     `json:"weatherCode"`
			WeatherDesc    string  `json:"weatherDesc"`
			WeatherIconURL string  `json:"weatherIconUrl"`
			Temp           int     `json:"temp"`
			Chanceofrain   int     `json:"chanceofrain"`
			RainInInches   float64 `json:"rainInInches"`
		} `json:"night"`
	} `json:"forecastOverview"`
	HighAtStr           interface{} `json:"highAtStr"`
	LoAtStr             interface{} `json:"loAtStr"`
	TimeZoneID          string      `json:"timeZoneId"`
	TimeFormat          string      `json:"timeFormat"`
	BarometerUnits      string      `json:"barometerUnits"`
	WindUnits           string      `json:"windUnits"`
	RainUnits           string      `json:"rainUnits"`
	TempUnits           string      `json:"tempUnits"`
	TemperatureFeelLike string      `json:"temperatureFeelLike"`
	Temperature         string      `json:"temperature"`
	HiTemp              string      `json:"hiTemp"`
	HiTempDate          int64       `json:"hiTempDate"`
	LoTemp              string      `json:"loTemp"`
	LoTempDate          int64       `json:"loTempDate"`
	Wind                string      `json:"wind"`
	Gust                string      `json:"gust"`
	GustAt              int64       `json:"gustAt"`
	Humidity            string      `json:"humidity"`
	Rain                string      `json:"rain"`
	SeasonalRain        string      `json:"seasonalRain"`
	Barometer           string      `json:"barometer"`
	BarometerTrend      string      `json:"barometerTrend"`
	LastReceived        int64       `json:"lastReceived"`
	SystemLocation      string      `json:"systemLocation"`
	AqsLocation         interface{} `json:"aqsLocation"`
	AqsLastReceived     interface{} `json:"aqsLastReceived"`
	ThwIndex            string      `json:"thwIndex"`
	ThswIndex           string      `json:"thswIndex"`
	Aqi                 interface{} `json:"aqi"`
	AqiString           interface{} `json:"aqiString"`
	AqiScheme           interface{} `json:"aqiScheme"`
	NoAccess            interface{} `json:"noAccess"`
}
