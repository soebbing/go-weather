package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const geolocationApiUrl = "https://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s"

type Location struct {
	Name    string  `json:"name"`
	Country string  `json:"country"`
	State   string  `json:"state"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

func GeoLocation(city, appId string) (*Location, error) {
	url := fmt.Sprintf(geolocationApiUrl, city, appId)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "GoWeather-API-Client")

	apiClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	res, getErr := apiClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintln("Status: ", res.StatusCode, string(body)))
	}

	var locations []*Location
	jsonErr := json.Unmarshal(body, &locations)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if len(locations) == 0 {
		return nil, errors.New("Not found")
	}

	return locations[0], nil
}
