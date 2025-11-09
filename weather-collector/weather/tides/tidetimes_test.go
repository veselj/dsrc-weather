package tides

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Scrape(t *testing.T) {
	res, err := Scrape()
	assert.Nil(t, err)
	assert.Greater(t, len(res), 0, "Expected at least one tide entry")
	fmt.Printf("%v", res)
}
