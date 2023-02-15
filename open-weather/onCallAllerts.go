package open_weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var URL = "https://api.openweathermap.org/data/3.0/onecall?%s"

type OneCallAlertData struct {
	Event       string `json:"event"`
	Start       int    `json:"start"`
	End         int    `json:"end"`
	Description string `json:"description"`
}

type OneCallData struct {
	Latitude  float64            `json:"lat"`
	Longitude float64            `json:"lon"`
	Timezone  string             `json:"timezone"`
	Alerts    []OneCallAlertData `json:"alerts,omitempty"`
	Key       string
	Excludes  string
	client    *http.Client
}

func (o *OneCallData) OneCallUsingCoordinates(lat, lon float64) error {
	url := fmt.Sprintf(fmt.Sprintf(URL, "appid=%s&lat=%f&lon=%f"), ApiKey, lat, lon)
	response, err := o.client.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	res := json.NewDecoder(response.Body).Decode(&o)
	return res
}
