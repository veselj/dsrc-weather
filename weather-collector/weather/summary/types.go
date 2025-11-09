package summary

// Sensor Type Id
const (
	Temperature = 7
	// DewPoint    = 10
	// WindChill        = 11
	// HeatIndex        = 12
	Humidity      = 14
	WindSpeed     = 15
	WindDirection = 17
	// RainStorm        = 20
	// RainTotal60Min   = 20
	// RainRate         = 22
	Barometer         = 26
	WinSpeed2MinAvg   = 53
	WinSpeedHigh2Min  = 54
	WinSpeed10MinAvg  = 55
	WinSpeedHigh10Min = 56
	// WetBulb          = 71
)

type WeatherSummary struct {
	OwnerName            interface{}       `json:"ownerName"`
	LastReceived         int64             `json:"lastReceived"` // in epoch milliseconds
	CurrConditionValues  []Value           `json:"currConditionValues"`
	HighLowValues        []Value           `json:"highLowValues"`
	AggregatedValues     []AggregatedValue `json:"aggregatedValues"`
	TimeSeriesValues     AdditionalData    `json:"timeSeriesValues"`
	TimeSeriesWeekValues AdditionalData    `json:"timeSeriesWeekValues"`
	AdditionalData       AdditionalData    `json:"additionalData"`
}

type AdditionalData struct {
}

type AggregatedValue struct {
	SensorDataName        string          `json:"sensorDataName"`
	RawValues             RawValues       `json:"rawValues"`
	ConvertedValues       ConvertedValues `json:"convertedValues"`
	AssocSensorDataTypeID interface{}     `json:"assocSensorDataTypeId"`
	UnitLabel             string          `json:"unitLabel"`
}

type ConvertedValues struct {
	Month string `json:"MONTH"`
	Year  string `json:"YEAR"`
	Day   string `json:"DAY"`
}

type RawValues struct {
	Month float64 `json:"MONTH"`
	Year  float64 `json:"YEAR"`
	Day   float64 `json:"DAY"`
}

type Value struct {
	SensorDataTypeID      *int64      `json:"sensorDataTypeId"`
	SensorDataName        string      `json:"sensorDataName"`
	DisplayName           *string     `json:"displayName"`
	ReportedValue         *float64    `json:"reportedValue"`
	Value                 *float64    `json:"value"`
	ConvertedValue        *string     `json:"convertedValue"`
	DepthLabel            interface{} `json:"depthLabel"`
	Category              *Category   `json:"category"`
	AssocSensorDataTypeID *int64      `json:"assocSensorDataTypeId"`
	SortOrder             *int64      `json:"sortOrder"`
	UnitLabel             *string     `json:"unitLabel"`
}

type Category string

const (
	High  Category = "high"
	Low   Category = "low"
	Main  Category = "main"
	Other Category = "other"
)
