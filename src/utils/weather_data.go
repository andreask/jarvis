package utils

import (
	"time"
)

type Timestamp struct {
	time.Time
}

type JCurrentWeather struct {
	Sunrise			*Timestamp 	`json:"sunrise"`
	Sunset			*Timestamp 	`json:"sunset"`
	Temp      		float64 	`json:"temp"`
	TempMin  		float64 	`json:"temp_min"`
	TempMax  		float64 	`json:"temp_max"`
	Pressure  		float64		`json:"pressure"`
	Humidity  		int 		`json:"humidity"`
	Weather			string 		`json:"weather"`
	Description		string 		`json:"description"`
	Icon			string 		`json:"icon"`
}

type JForecastWeather struct {
	TempDay			float64	`json:"temp_day"`
	TempMin			float64	`json:"temp_min"`
	TempMax			float64	`json:"temp_max"`
	TempNight		float64	`json:"temp_night"`
	Pressure  		float64 `json:"pressure"`
	Humidity  		int 	`json:"humidity"`
	Weather			string 	`json:"weather"`
	Description		string 	`json:"description"`
	Icon			string 	`json:"icon"`
}

type JWeather struct {
	JCurrentWeather JCurrentWeather				`json:"current_weather"`
	JForecastWeatherList [5]JForecastWeather	`json:"forecast_weather"`
}
