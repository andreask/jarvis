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
	ID				int 		`json:"id"`
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
	ID				int 	`json:"id"`
}

type JWeather struct {
	Updated					*Timestamp			`json:"updated"`
	JCurrentWeather 		JCurrentWeather		`json:"current_weather"`
	JForecastWeatherList 	[]JForecastWeather	`json:"forecast_weather"`
}

type JStationModule struct {
	ID string
	Name string
	TimeStamp int
	Values JWeatherValues
}

type JWeatherValues struct {
	Temperature         float32
	Humidity            int32
	CO2                 int32
	Noise               int32
	Pressure            float32
}
