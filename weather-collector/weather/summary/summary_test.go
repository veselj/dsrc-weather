package summary

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Summary(t *testing.T) {
	w, err := GetSummaryWeather()
	assert.NoError(t, err)
	fmt.Printf("%+v", w)
}
