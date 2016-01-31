package main

import (
	"transportation/api"
	"fmt"
)

func main() {
	siteId := api.SearchPlace("Tekniska högskolan")
	for _,info := range api.GetDepartures(siteId, "Metros") {
		fmt.Printf("%s\t%s %s - %s\n", info.DisplayTime, info.LineNumber, info.Destination, info.ExpectedDateTime)
	}

	siteId = api.SearchPlace("Ruddammen")
	for _,info := range api.GetDepartures(siteId, "Buses") {
		fmt.Printf("%s\t%s %s - %s\n", info.DisplayTime, info.LineNumber, info.Destination, info.ExpectedDateTime)
	}
}
