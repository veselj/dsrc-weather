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
