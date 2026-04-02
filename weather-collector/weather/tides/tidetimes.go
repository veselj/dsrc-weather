package tides

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	LowTide = iota
	HighTide
)

type Tide struct {
	Type   int
	Time   int64 // Unix timestamp
	Height float64
}

func Scrape() ([]Tide, error) {
	resp, err := http.Get("https://www.tidetimes.org.uk/dundee-tide-times")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	tides := []Tide{}
	
	tidesDiv := doc.Find("div#tides")
	if tidesDiv.Length() > 0 {
		table := tidesDiv.Find("table").First()
		table.Find("tr").Each(func(i int, row *goquery.Selection) {
			// Skip rows with class "vis0"
			if row.HasClass("vis0") {
				return
			}
			cells := row.Find("td")
			if cells.Length() == 3 {
				tideType := parseTideType(strings.TrimSpace(cells.Eq(0).Text()))
				if tideType == -1 {
					return
				}
				tideTime := parseTideTime(strings.TrimSpace(cells.Eq(1).Find("span").Text()))
				height := parseTideHeight(strings.TrimSpace(cells.Eq(2).Text()))
				tides = append(tides, Tide{
					Type:   tideType,
					Time:   tideTime,
					Height: height,
				})
			}
		})
	}

	return tides, nil
}

func parseTideType(tideType string) int {
	if strings.Contains(tideType, "High") {
		return HighTide
	} else if strings.Contains(tideType, "Low") {
		return LowTide
	}
	return -1
}

func parseTideTime(tideTime string) int64 {
	loc, err := time.LoadLocation("Europe/London")
	if err != nil {
		loc = time.UTC
	}
	now := time.Now().In(loc)
	parsed, err := time.ParseInLocation("15:04", tideTime, loc)
	if err != nil {
		return time.Time{}.Unix()
	}
	// Combine today's date with parsed hour and minute
	return time.Date(
		now.Year(), now.Month(), now.Day(),
		parsed.Hour(), parsed.Minute(), 0, 0, loc,
	).Unix()
}

func parseTideHeight(tideHeight string) float64 {
	heightStr := strings.TrimSuffix(tideHeight, "m")
	height, _ := strconv.ParseFloat(strings.TrimSpace(heightStr), 64)
	return height
}
