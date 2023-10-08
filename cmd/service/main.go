package main

import (
	"os"

	"github.com/abkhan/simpleweather/internal/server"
	"github.com/abkhan/simpleweather/internal/weatherapi"
)

func main() {

	// Get apiKey from env
	url := os.Getenv("OPENWEATHER_URL") // example: https://api.openweathermap.org/data/2.5/weather
	if url == "" {
		url = "https://api.openweathermap.org/data/2.5/weather"
	}
	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	// first create weatherApi
	wapi := weatherapi.New(url, apiKey)

	// then create and start server
	wserver := server.New(wapi)
	wserver.Start()
}
