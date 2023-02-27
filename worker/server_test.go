package worker

import (
	open_weather "github.com/nicknikandish/openweatherApp/open-weather"
	"net/http"
	"reflect"
	"testing"
)

func TestWeatherHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WeatherHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_getAlerts(t *testing.T) {
	type args struct {
		wd open_weather.CityWeatherData
	}
	tests := []struct {
		name    string
		args    args
		want    *owm.OneCallData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAlerts(tt.args.wd)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAlerts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAlerts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCurrent(t *testing.T) {
	type args struct {
		lat float64
		lon float64
	}
	tests := []struct {
		name    string
		args    args
		want    *owm.CityWeatherData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCurrent(tt.args.lat, tt.args.lon)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCurrent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCurrent() got = %v, want %v", got, tt.want)
			}
		})
	}
}
