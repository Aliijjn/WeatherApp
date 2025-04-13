package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getLocation() (Location, error) {
	const apiURL = "https://ipinfo.io/json"

	response, err := http.Get(apiURL)
	if err != nil {
		return Location{}, fmt.Errorf("Error with API call: %d", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Location{}, fmt.Errorf("Bad response code: %d", response.StatusCode)
	}
	var output struct {
		City   string `json:"city"`
		Region string `json:"region"`
		Str    string `json:"loc"`
	}
	err = json.NewDecoder(response.Body).Decode(&output)
	if err != nil {
		return Location{}, fmt.Errorf("Error parsing JSON response: %d", err)
	}
	split := strings.Split(output.Str, ",")
	if len(split) != 2 {
		return Location{}, fmt.Errorf("Invalid location format: '%s'", output.Str)
	}
	location := Location{output.City, output.Region, split[0], split[1]}
	return location, nil
}

func getWeather(location Location) (WeatherInfo, error) {
	apiURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&hourly=temperature_2m,apparent_temperature,precipitation,wind_speed_10m,wind_direction_10m", location.Latitude, location.Longitude)

	// API call
	response, err := http.Get(apiURL)
	if err != nil {
		return WeatherInfo{}, fmt.Errorf("Error with API call: %d", err)
	}
	defer response.Body.Close()

	// Parse response
	if response.StatusCode != http.StatusOK {
		return WeatherInfo{}, fmt.Errorf("Bad response code: %d", response.StatusCode)
	}
	var weather WeatherResponse
	err = json.NewDecoder(response.Body).Decode(&weather)
	if err != nil {
		return WeatherInfo{}, fmt.Errorf("Error parsing JSON response: %d", err)
	}
	return ToWeatherInfo(weather, location), nil
}

func main() {
	location, err := getLocation()
	if err != nil {
		log.Fatalf("Could not get location: '%v'", err)
	}
	weather, err := getWeather(location)
	if err != nil {
		log.Fatalf("Could not get local weather: '%v'", err)
	}
	fmt.Println(weather.ToString())
}