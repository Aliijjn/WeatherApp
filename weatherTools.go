package main

import (
	"fmt"
	"math"
)

type Location struct {
	City      string
	Region    string
	Latitude  string
	Longitude string
}

type HourlyWeather struct {
	Day           string
	TimeOfDay     string
	Temperature   float64
	ApTemperature float64
	Precipitation float64
	WindSpeed     float64
	WindDirection string
}

func (weather HourlyWeather) ToString() string {
	return fmt.Sprintf("%s  │ %5.1f°C  │ %5.1f°C  │ %5.1fmm  │ %3.0fkm/h %-2s", 
		weather.TimeOfDay, 
		weather.Temperature, 
		weather.ApTemperature, 
		weather.Precipitation, 
		weather.WindSpeed, 
		weather.WindDirection)
}

type WeatherInfo struct {
	Weather  []HourlyWeather
	Location Location
}

func (weatherInfo WeatherInfo) ToString() string {
	result := ""
	previousDay := ""
	for _, weather := range weatherInfo.Weather {
		if previousDay != weather.Day {
			if previousDay != "" {
				result += "└────────────────────────────────────────────────────────┘\n"
			} 
			result += "╔════════════════════════════════════════════════════════╗\n"
			result += fmt.Sprintf("║  %s - %-45.43s ║\n", weather.Day, fmt.Sprintf("%s, %s", weatherInfo.Location.City, weatherInfo.Location.Region))
			result += "╟─────────┬──────────┬──────────┬──────────┬─────────────╢\n"
			result += "║  Time   │   Temp   │ Feels as │   Rain   │    Wind     ║\n"
			result += "╟═════════╧══════════╧══════════╧══════════╧═════════════╢\n"
		}
		result += ("│  " + weather.ToString() + "  │\n")
		previousDay = weather.Day
	}
	result += "└────────────────────────────────────────────────────────┘\n"
	return result
} 

func reformatTime(time string) (string, string) {
	
	date := time[5:10]
	timeOfDay := time[11:]
	return date, timeOfDay
}

func reformatWindDir(dir float64) string {
	directions := []struct {
		angle float64;
		label string
	}{
		{0,   "N"},
		{45,  "NE"},
		{90,  "E"},
		{135, "SE"},
		{180, "S"},
		{225, "SW"},
		{270, "W"},
		{315, "NW"},
	}
	dir = math.Mod(dir + 22.5, 360)
	index := int(dir / 45)
	return directions[index].label
}

func ToWeatherInfo(weather WeatherResponse, location Location) WeatherInfo {
	result := WeatherInfo{
		Location: location,
	}
	for i := range weather.Hourly.Time {
		day, timeOfDay := reformatTime(weather.Hourly.Time[i])
		result.Weather = append(result.Weather, HourlyWeather{
			Day:           day,
			TimeOfDay:     timeOfDay,
			Temperature:   weather.Hourly.Temperature2m[i],
			ApTemperature: weather.Hourly.ApTemperature[i],
			Precipitation: weather.Hourly.Precipitation[i],
			WindSpeed:     weather.Hourly.WindSpeed[i],
			WindDirection: reformatWindDir(weather.Hourly.WindDirection[i]),
		})
	}
	return result
}

type WeatherResponse struct {
	Hourly struct {
		Time                []string  `json:"time"`
		Temperature2m       []float64 `json:"temperature_2m"`
		ApTemperature       []float64 `json:"apparent_temperature"`
		Precipitation       []float64 `json:"precipitation"`
		WindSpeed           []float64 `json:"wind_speed_10m"`
		WindDirection       []float64 `json:"wind_direction_10m"`
	} `json:"hourly"`
}