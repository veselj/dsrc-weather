package station

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	rtycWeatherLink = "https://www.weatherlink.com/embeddablePage/getData/477837b179b94d58b123a4c127c40c50"
)

func GetWeather() (*WeatherData, error) {
	return fetchWeatherData(rtycWeatherLink)
}

func fetchWeatherData(url string) (*WeatherData, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("User-Agent", "PostmanRuntime/7.43.4")
	req.Header.Set("Accept", "application/json")

	// Make request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON
	var weatherResp WeatherData
	if err := json.Unmarshal(body, &weatherResp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON body: %w", err)
	}

	return &weatherResp, nil
}
