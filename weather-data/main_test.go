package main

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_retrieveSamples(t *testing.T) {
	ctx := context.Background()
	samples := retrieveSamples(ctx, time.Now().Add(-2*time.Hour).Unix())
	samplesJSON, err := json.MarshalIndent(samples, "", "  ")
	assert.NoError(t, err)
	fmt.Println(string(samplesJSON))
}
