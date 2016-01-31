package api

import (
	"net/http"
	"encoding/json"
	"os"
	"net/url"
	"strings"
)

const (
	placeInformationUrl = "http://api.sl.se/api2/typeahead.json"
)

type Place struct {
	Name 	string `json:"Name"`
	SiteID 	string `json:"SiteId"`
	Type	string `json:"Type"`
}

type PlaceSearchResults struct {
	StatusCode		int `json:"StatusCode"`
	Message			string `json:"Message"`
	ExecutionTime	int `json:"ExecutionTime"`
	Places	 		[]*Place `json:"ResponseData"`
}

func panic_error(err error) {
	if err != nil {
		panic(err)
	}
}

func SearchPlace(placeName string) string {
	url := placeInformationUrl + "?key=" + os.Getenv("PLACE_INFO_API_KEY") + "&searchString=" + url.QueryEscape(placeName)

	resp, err := http.Get(url)
	panic_error(err)
	defer resp.Body.Close()

	var data PlaceSearchResults

	err = json.NewDecoder(resp.Body).Decode(&data)
	panic_error(err)

	for _, place := range data.Places {
		if strings.HasPrefix(place.Name, placeName) {
			return place.SiteID
		}
	}

	return ""
}