package store

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getHourlyBucketsNow(t *testing.T) {
	starts := GetHourlyBucketStarts(time.Now().Unix())
	assert.Equal(t, 1, len(starts))
}

func Test_getHourlyBucketsOneHourBefore(t *testing.T) {
	starts := GetHourlyBucketStarts(time.Now().Add(-time.Hour).Unix())
	assert.Equal(t, 2, len(starts))
}

func Test_getHourlyBucketsTwoHourBefore(t *testing.T) {
	starts := GetHourlyBucketStarts(time.Now().Add(-2 * time.Hour).Unix())
	assert.Equal(t, 3, len(starts))
}
