package tides

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Tide struct {
	Type   string
	Time   string
	Height string
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
	doc.Find("table").EachWithBreak(func(_ int, table *goquery.Selection) bool {
		header := table.Find("thead tr th").First().Text()
		if strings.Contains(header, "Tide Times") {
			table.Find("tr").Each(func(_ int, row *goquery.Selection) {
				cells := row.Find("td")
				if cells.Length() == 3 {
					tides = append(tides, Tide{
						Type:   strings.TrimSpace(cells.Eq(0).Text()),
						Time:   strings.TrimSpace(cells.Eq(1).Find("span").Text()),
						Height: strings.TrimSpace(cells.Eq(2).Text()),
					})
				}
			})
			return false // Stop after finding the correct table
		}
		return true
	})

	return tides, nil
}
