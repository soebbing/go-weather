package main

import (
	"GoWeather/weather/api"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Benutzung:")
		fmt.Println("weather (Stadt,Landeskürzel)")
		os.Exit(1)
	}

	appId := os.Getenv("OPENWEATHER_APP_ID")

	if appId == "" {
		fmt.Println("ENV Variable mit OpenWeather AppId konnte nicht gefunden werden. Setzen mit:")
		fmt.Println("export OPENWEATHER_APP_ID=(APP_ID)")
		os.Exit(1)
	}

	location, err := api.GeoLocation(os.Args[1], appId)

	if err != nil {
		log.Fatal(err)
	}

	weather, err := api.GetWeather(*location, appId)

	fmt.Printf(
		"%s: %s %s bei %.1f°\n",
		weather.Name,
		weather.Current[0].Description,
		weather.Current[0].Icon,
		weather.Main.Temp,
	)
}
