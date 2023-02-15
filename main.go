package main

import (
	"fmt"
	owm "github.com/nicknikandish/openweatherApp/open-weather"
	"net/http"
	"os"
	"strconv"
)

// getCurrent gets the current weather for the provided location in
// the units provided.
func getCurrent(lat, lon float64) (*owm.CityWeatherData, error) {
	w, err := owm.NewCurrent(os.Getenv("OWM_APP_ID")) // Create the instance with the given unit
	if err != nil {
		return nil, err
	}
	w.CurrentWeatherFromCoordinates(lat, lon) // Get the actual data for the given location
	return w, nil
}

func getAlerts(wd owm.CityWeatherData) (*owm.OneCallData, error) {
	c, err := owm.NewOneCall(os.Getenv("OWM_APP_ID"))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = c.OneCallUsingCoordinates(wd.Pos.Latitude, wd.Pos.Longitude)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return c, nil

}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	var cor []string
	var err error
	var lat, lon float64
	for _, v := range r.URL.Query() {
		cor = append(cor, v...)
	}
	if lat, err = strconv.ParseFloat(cor[0], 64); err != nil {
		fmt.Println(err)
		return
	}

	if lon, err = strconv.ParseFloat(cor[1], 64); err != nil {
		fmt.Println(err)
		return
	}
	wd, err := getCurrent(lat, lon)
	if err != nil {
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}
	c, err := getAlerts(*wd)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "<h1 align=\"center\">City: %v</h1>", wd.Name)
	if len(wd.Weather) > 0 {
		fmt.Fprintf(w, "<h3 align=\"center\">Weather desscription: %v</h3>", wd.Weather[0].Description)
	}
	fmt.Fprintf(w, "<h3 align=\"center\">Temperature: %v</h3>", wd.Main.Temp)
	fmt.Fprintf(w, "<h3 align=\"center\">TempMax: %v</h3>", wd.Main.TempMax)
	fmt.Fprintf(w, "<h3 align=\"center\">TempMin: %v</h3>", wd.Main.TempMin)
	if wd.Main.Temp > 50.00 && wd.Main.Temp < 86.00 {
		fmt.Fprintf(w, "<h3 align=\"center\">Weather conditions: %v</h3>", "Moderate")
	} else if wd.Main.Temp <= 45.00 {
		fmt.Fprintf(w, "<h3 align=\"center\">Weather conditions: %v</h3>", "Cold")
	} else if wd.Main.Temp >= 86.00 {
		fmt.Fprintf(w, "<h3 align=\"center\">Weather conditions: %v</h3>", "Hot")
	}
	fmt.Fprintf(w, "<h3 align=\"center\">Humidity: %v</h3>", wd.Main.Humidity)
	fmt.Fprintf(w, "<h3 align=\"center\">Snow: %v</h3>", wd.Snow)
	fmt.Fprintf(w, "<h3 align=\"center\">Rain: %v</h3>", wd.Rain)
	fmt.Fprintf(w, "<h3 align=\"center\">Visibility: %v</h3>", wd.Visibility)
	fmt.Fprintf(w, "<h3 align=\"center\">Alerts: %v</h3>", c.Alerts)

}

func main() {

	http.HandleFunc("/coordination", weatherHandler)

	http.ListenAndServe(":8000", nil)
}
