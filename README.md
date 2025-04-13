# WeatherApp
Welcome to WeatherApp. This is my first (but certainly not the last) attempt at making an application in Go!

## What I learned from this project
- The basics of Go
- How to interact with APIs
- How to cleanly parse JSON files in Go

## How to run the program
If you have Go installed, simply run `go run main.go weatherTools.go` in your terminal

## How this program works
1. The program fetches your general location using [IPInfo's API](https://ipinfo.io/json)
2. It then uses the latitude and longitude to find your region's weather using [Open-metio's API](https://api.open-meteo.com/v1/forecast?latitude=52.3&longitude=4.9&hourly=temperature_2m,apparent_temperature,precipitation,wind_speed_10m,wind_direction_10m)
3. It parses the JSON into an array of `HourlyWeather`
4. It prints this array with some formatting

## Output:
![image](https://github.com/user-attachments/assets/6a11eb8d-3cb5-470b-ab10-895297288f46)

```
╔════════════════════════════════════════════════════════╗
║  04-13 - Utrecht, Utrecht                              ║
╟─────────┬──────────┬──────────┬──────────┬─────────────╢
║  Time   │   Temp   │ Feels as │   Rain   │    Wind     ║
╟═════════╧══════════╧══════════╧══════════╧═════════════╢
│  00:00  │  16.9°C  │  14.2°C  │   0.0mm  │  13km/h S   │
│  01:00  │  17.0°C  │  14.3°C  │   0.0mm  │  16km/h S   │
│  02:00  │  13.9°C  │  12.9°C  │   0.4mm  │  10km/h S   │
│  03:00  │  14.2°C  │  12.2°C  │   0.4mm  │  18km/h S   │
│  04:00  │  13.9°C  │  12.8°C  │   0.7mm  │  12km/h S   │
│  05:00  │  14.0°C  │  12.0°C  │   0.2mm  │  19km/h SW  │
│  06:00  │  13.9°C  │  12.6°C  │   0.4mm  │  14km/h SW  │
│  07:00  │  14.3°C  │  12.7°C  │   0.1mm  │  17km/h SW  │
│  08:00  │  14.4°C  │  12.3°C  │   0.1mm  │  19km/h SW  │
│  09:00  │  13.9°C  │  11.5°C  │   0.0mm  │  19km/h SW  │
│  10:00  │  14.3°C  │  12.2°C  │   0.6mm  │  19km/h W   │
│  11:00  │  14.2°C  │  13.0°C  │   0.0mm  │  10km/h W   │
│  12:00  │  15.4°C  │  13.7°C  │   0.0mm  │  14km/h SW  │
│  13:00  │  14.2°C  │  12.9°C  │   0.0mm  │  12km/h W   │
│  14:00  │  15.0°C  │  13.8°C  │   0.0mm  │  12km/h W   │
│  15:00  │  14.1°C  │  13.2°C  │   0.0mm  │  10km/h W   │
│  16:00  │  15.3°C  │  13.7°C  │   0.0mm  │  12km/h SW  │
│  17:00  │  14.9°C  │  13.0°C  │   0.1mm  │  10km/h W   │
│  18:00  │  14.3°C  │  12.3°C  │   0.1mm  │  13km/h W   │
│  19:00  │  13.5°C  │  11.2°C  │   0.0mm  │  13km/h W   │
│  20:00  │  12.6°C  │  11.0°C  │   0.0mm  │   7km/h SW  │
│  21:00  │  12.0°C  │  10.3°C  │   0.0mm  │  10km/h S   │
│  22:00  │  11.4°C  │   9.5°C  │   0.0mm  │  11km/h SW  │
│  23:00  │  11.0°C  │   8.7°C  │   0.0mm  │  12km/h S   │
└────────────────────────────────────────────────────────┘
