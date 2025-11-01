package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"time"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/veselj/dsrc-weather/internal/record"
	"github.com/veselj/dsrc-weather/internal/store"
)

func handler(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {

	from := time.Now().Add(-time.Hour * 3)

    if hoursStr, ok := req.QueryStringParameters["hours"]; ok {
        hours, err := strconv.ParseInt(hoursStr, 10, 64)
        if err == nil {
            from = time.Now().Add(-time.Hour * time.Duration(hours))
        }
    }

	samples := retrieveSamples(ctx, from.Unix())
	jsonBody, err := json.Marshal(samples)
	if err != nil {
		jsonBody = []byte(err.Error())
	}

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	gz.Write(jsonBody)
	gz.Close()

	// Base64 encode
	gzippedBase64 := base64.StdEncoding.EncodeToString(b.Bytes())

	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":     "application/json",
			"Content-Encoding": "gzip",
		},
		Body:            gzippedBase64,
		IsBase64Encoded: true,
	}, nil
}

func main() {
	lambda.Start(handler)
}

func retrieveSamples(ctx context.Context, fromUnix int64) []record.Sample {
	dynamo := store.NewDynamoClient()
	startTimes := store.GetHourlyBucketStarts(fromUnix)
	retrieved := []record.Sample{}
	for _, startTime := range startTimes {
		s := dynamo.Samples(ctx, startTime)
		retrieved = append(retrieved, s...)
	}
	return retrieved
}
