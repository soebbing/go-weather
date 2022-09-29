package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const weatherApiUrl = "https://api.openweathermap.org/data/2.5/weather?units=metric&lang=%s&lat=%.7f&lon=%.7f&exclude=minutely&appid=%s"

var iconMap = map[string]string{
	"01d": "☀", // clear sky
	"02d": "🌤", // few clouds
	"03d": "☁", // scattered clouds
	"04d": "🌥", // broken clouds
	"09d": "🌧", // shower rain
	"10d": "🌦", // rain
	"11d": "🌩", // thunderstorm
	"13d": "🌨", // snow
	"50d": "🌫", // mist

	"01n": "☀", // clear sky
	"02n": "🌤", // few clouds
	"03n": "☁", // scattered clouds
	"04n": "🌥", // broken clouds
	"09n": "🌧", // shower rain
	"10n": "🌦", // rain
	"11n": "🌩", // thunderstorm
	"13n": "🌨", // snow
	"50n": "🌫", // mist
}

type Weather struct {
	Coordinates coordinates `json:"coord"`
	Current     []current   `json:"weather"`
	Base        string      `json:"base"`
	Main        main        `json:"main"`
	Visibility  int64       `json:"visibility"`
	Wind        wind        `json:"wind"`
	Clouds      clouds      `json:"clouds"`
	Sys         sys         `json:"sys"`
	DateTime    int         `json:"dt"`
	Timezone    int         `json:"timezone"`
	Cod         int         `json:"cod"`
	Name        string      `json:"name"`
}

type coordinates struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

type current struct {
	Id          int16  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type wind struct {
	Speed  float32 `json:"speed"`
	Gust   float32 `json:"gust"`
	Degree int16   `json:"deg"`
}

type clouds struct {
	All int8 `json:"all"`
}

type main struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  int16   `json:"pressure"`
	Humidity  int8    `json:"humidity"`
}

type sys struct {
	Type    int8   `json:"type"`
	Id      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int32  `json:"sunrise"`
	Sunset  int32  `json:"sunset"`
}

func GetWeather(city Location, appId string) (*Weather, error) {
	url := fmt.Sprintf(weatherApiUrl, "de", city.Lat, city.Lon, appId)

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

	var weather *Weather
	jsonErr := json.Unmarshal(body, &weather)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	icon, ok := iconMap[weather.Current[0].Icon]
	if ok == false {
		panic(fmt.Sprintln(weather.Current[0].Icon, "fehlt"))
	}

	weather.Current[0].Icon = icon

	return weather, nil
}
