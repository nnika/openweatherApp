package open_weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var ApiKey = os.Getenv("OWM_APP_ID")

const ApiUrl string = "api.openweathermap.org"

// ConditionData holds data structure for weather conditions information.
type ConditionData struct {
	ID          int
	Description string
}

type Main struct {
	Temp     float64 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

// Coordinates struct holds longitude and latitude data in returned
// JSON or as parameter data for requests using longitude and latitude.
type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type Rain struct {
	ThreeH float64 `json:"3h,omitempty"`
}

type Snow struct {
	ThreeH float64 `json:"3h,omitempty"`
}

type City struct {
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	Coord   Coordinates `json:"coord"`
	Country string      `json:"country"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

type CityWeatherData struct {
	Pos        Coordinates        `json:"coord"`
	Weather    []Weather          `json:"weather"`
	Main       Main               `json:"main"`
	Visibility int                `json:"visibility"`
	Wind       Wind               `json:"wind"`
	Rain       Rain               `json:"rain"`
	Snow       Snow               `json:"snow"`
	ID         int                `json:"id"`
	Name       string             `json:"name"`
	Alerts     []OneCallAlertData `json:"alerts,omitempty"`
	Key        string
	client     *http.Client
}

func setKey(key string) (string, error) {
	return key, nil
}

func NewCurrent(key string) (*CityWeatherData, error) {
	c := &CityWeatherData{
		client: http.DefaultClient,
	}

	var err error
	c.Key, err = setKey(key)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func NewOneCall(key string) (*OneCallData, error) {
	o := &OneCallData{
		client: http.DefaultClient,
	}

	var err error
	o.Key, err = setKey(key)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (w *CityWeatherData) CurrentWeatherFromCoordinates(lat, lon float64) error {
	response, err := w.client.Get(fmt.Sprintf("https://%s/data/2.5/weather?lat=%f&lon=%f&units=imperial&APPID=%s", ApiUrl, lat, lon, ApiKey))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		return err
	}

	if err = json.NewDecoder(response.Body).Decode(&w); err != nil {
		return err
	}

	return nil
}
