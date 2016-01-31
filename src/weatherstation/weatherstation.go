package main

import (
	"fmt"
	"time"
	"os"
	"github.com/andreask/netatmo-api-go"
	"encoding/json"
	"utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	n, err := netatmo.NewClient(netatmo.Config{
		ClientID:     os.Getenv("NETATMO_CLIENTID"),
		ClientSecret: os.Getenv("NETATMO_CLIENTSECRET"),
		Username:     os.Getenv("NETATMO_USERNAME"),
		Password:     os.Getenv("NETATMO_PASSWORD"),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dc, err := n.GetDeviceCollection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, station := range dc.Stations() {
		if station.ID == "70:ee:50:05:66:2a" {
			modules := []utils.JStationModule{}
			f, err := os.Create(os.Getenv("GOPATH") + "/output/weather_station.json")

			check(err)

			fmt.Printf("Station : %s (%s)\n", station.StationName, station.ID)

			for _, module := range station.Modules() {

				fmt.Printf("\tModule : %s (%s)\n", module.ModuleName, module.ID)

				jsonModule := utils.JStationModule {
					ID: module.ID,
					Name: module.ModuleName,
					TimeStamp: int(module.DashboardData.LastMeasure),
					Values: utils.JWeatherValues {
						CO2: module.DashboardData.CO2,
						Humidity: module.DashboardData.Humidity,
						Temperature: module.DashboardData.Temperature,
						Noise: module.DashboardData.Noise,
						Pressure: module.DashboardData.Pressure } }

				ts, data := module.Data()
				for dataType, value := range data {
					switch value.(type) {
					case int32:
						fmt.Printf("\t\t%s : %d (%s)\n", dataType, value.(int32), time.Unix(int64(ts), 0))
					case float32:
						fmt.Printf("\t\t%s : %.1f (%s)\n", dataType, value.(float32), time.Unix(int64(ts), 0))
					}
				}

				modules = append(modules, jsonModule)
			}

			json, err := json.Marshal(modules)

			if err == nil {
				f.Write(json)
			}

			f.Close()
		}
	}
}