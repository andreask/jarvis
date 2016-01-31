package api
import (
	"os"
	"net/http"
	"encoding/json"
)

const (
	realTimeInformationUrl = "http://api.sl.se/api2/realtimedepartures.json"
)

type TraficInfo struct {
	DisplayTime			string `json:"DisplayTIme"`
	Destination			string `json:"Destination"`
	LineNumber			string `json:"LineNumber"`
	TransportMode		string `json:"TransportMode"`
	GroupOfLine			string `json:"GroupOfLine"`
	TimeTabledDateTime 	string `json:"TimeTableDateTime"`
	ExpectedDateTime	string `json:"ExpectedDateTime"`
}

type ResponseData struct {
	LatestUpdate	string `json:"LatestUpdate"`
	DataAge			int `json:"DataAge"`
	Metros			[]*TraficInfo `json:"Metros"`
	Buses			[]*TraficInfo `json:"Buses"`
}

type TraficInfoSearchResults struct {
	StatusCode		int `json:"StatusCode"`
	Message			string `json:"Message"`
	ExecutionTime	int `json:"ExecutionTime"`
	ResponseData	ResponseData `json:"ResponseData"`
}

func GetDepartures(SiteId string, Type string) []*TraficInfo {
	url := realTimeInformationUrl + "?key=" + os.Getenv("REALTIME_INFO_API_KEY") + "&siteId=" + SiteId

	resp, err := http.Get(url)
	panic_error(err)
	defer resp.Body.Close()

	var data TraficInfoSearchResults

	err = json.NewDecoder(resp.Body).Decode(&data)
	panic_error(err)

	if (Type == "Metros") {
		return data.ResponseData.Metros
	}

	if (Type == "Buses") {
		return data.ResponseData.Buses
	}

	return nil
}