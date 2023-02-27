package open_weather

import (
	"reflect"
	"testing"
)

func TestCityWeatherData_CurrentWeatherFromCoordinates(t *testing.T) {
	type fields struct {
		Pos        Coordinates
		Weather    []Weather
		Main       Main
		Visibility int
		Wind       Wind
		Rain       Rain
		Snow       Snow
		ID         int
		Name       string
		Alerts     []OneCallAlertData
		Key        string
		client     *http.Client
	}
	type args struct {
		lat float64
		lon float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &CityWeatherData{
				Pos:        tt.fields.Pos,
				Weather:    tt.fields.Weather,
				Main:       tt.fields.Main,
				Visibility: tt.fields.Visibility,
				Wind:       tt.fields.Wind,
				Rain:       tt.fields.Rain,
				Snow:       tt.fields.Snow,
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				Alerts:     tt.fields.Alerts,
				Key:        tt.fields.Key,
				client:     tt.fields.client,
			}
			if err := w.CurrentWeatherFromCoordinates(tt.args.lat, tt.args.lon); (err != nil) != tt.wantErr {
				t.Errorf("CurrentWeatherFromCoordinates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCurrent(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    *CityWeatherData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCurrent(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCurrent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCurrent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("setKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("setKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
