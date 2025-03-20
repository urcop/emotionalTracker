package horoscope

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/urcop/emotionalTracker/domain/services"
)

const (
	horoscopeAPIURL = "https://horoscope-app-api.vercel.app/api/v1/get-horoscope/daily"
)

type HoroscopeService struct{}

type HoroscopeAPIResponse struct {
	Data struct {
		Date          string `json:"date"`
		HoroscopeData string `json:"horoscope_data"`
	} `json:"data"`
	Status  int  `json:"status"`
	Success bool `json:"success"`
}

func NewHoroscopeService() (*HoroscopeService, error) {
	return &HoroscopeService{}, nil
}

// GetDailyHoroscope fetches the daily horoscope for the given zodiac sign
func (h *HoroscopeService) GetDailyHoroscope(sign string) (*services.HoroscopeResponse, error) {
	// Build the URL with query parameters
	url := fmt.Sprintf("%s?sign=%s&day=TODAY", horoscopeAPIURL, sign)

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch horoscope: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	var horoscopeResp HoroscopeAPIResponse
	if err := json.Unmarshal(body, &horoscopeResp); err != nil {
		return nil, fmt.Errorf("failed to parse horoscope response: %w", err)
	}

	// Check if the request was successful
	if !horoscopeResp.Success || horoscopeResp.Status != 200 {
		return nil, fmt.Errorf("horoscope API returned error status: %d", horoscopeResp.Status)
	}

	// Return the horoscope
	return &services.HoroscopeResponse{
		Date:          horoscopeResp.Data.Date,
		HoroscopeData: horoscopeResp.Data.HoroscopeData,
		ZodiacSign:    sign,
	}, nil
}

// Close closes the resources held by the service
func (h *HoroscopeService) Close() error {
	return nil
}
