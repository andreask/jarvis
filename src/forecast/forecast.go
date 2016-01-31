package main

import (
	"log"
	owm "github.com/briandowns/openweathermap"
	"encoding/json"
	"os"
	"utils"
	"time"
)

func main() {
	jWeather := utils.JWeather{}

	f, err := os.Create(os.Getenv("GOPATH") + "/output/openweather.json")

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	defer f.Close()

	current, err := owm.NewCurrent("C", "sv")
	if err != nil {
		log.Fatalln(err)
	}

	forecast, err := owm.NewForecast("C", "sv")

	if err != nil {
		log.Fatalln(err)
	}

	err = forecast.DailyByID(2673730, 5)

	if err != nil {
		log.Fatalln(err)
	} else {
		for i, weather := range forecast.List {
			jWeather.JForecastWeatherList[i] = utils.JForecastWeather{
				TempDay: weather.Temp.Day,
				TempMin: weather.Temp.Min,
				TempMax: weather.Temp.Max,
				TempNight: weather.Temp.Night,
				Pressure: weather.Pressure,
				Humidity: weather.Humidity,
				Weather: weather.Weather[0].Main,
				Description: weather.Weather[0].Description,
				Icon: weather.Weather[0].Icon,
			}
		}
	}

	err = current.CurrentByID(2673730)
	if err != nil {
		log.Fatalln(err)
	} else {
		jWeather.JCurrentWeather = utils.JCurrentWeather{
			Sunrise: &utils.Timestamp { Time: time.Unix(int64(current.Sys.Sunrise), 0) },
			Sunset: &utils.Timestamp { Time: time.Unix(int64(current.Sys.Sunset), 0) },
			Temp: current.Main.Temp,
			TempMin: current.Main.TempMin,
			TempMax: current.Main.TempMax,
			Pressure: current.Main.Pressure,
			Humidity: current.Main.Humidity,
			Weather: current.Weather[0].Main,
			Description: current.Weather[0].Description,
			Icon: current.Weather[0].Icon,
		}

		byteArray, err := json.Marshal(jWeather)

		if err == nil {
			f.Write(byteArray)
		} else {
			log.Fatalln(err)
		}
	}
}
