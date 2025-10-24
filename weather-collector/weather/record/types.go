package record

type Sample struct {
	Wind        float64
	Direction   int
	Temperature float64
	FeelsLike   float64
	When        int64 // Unix timestamp
	Bucket      string
}
