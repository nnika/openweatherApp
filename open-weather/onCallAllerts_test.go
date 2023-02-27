package open_weather

import (
	"reflect"
	"testing"
)

func TestNewOneCall(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    *OneCallData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOneCall(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOneCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOneCall() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOneCallData_OneCallUsingCoordinates(t *testing.T) {
	type fields struct {
		Latitude  float64
		Longitude float64
		Timezone  string
		Alerts    []OneCallAlertData
		Key       string
		Excludes  string
		client    *http.Client
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
			o := &OneCallData{
				Latitude:  tt.fields.Latitude,
				Longitude: tt.fields.Longitude,
				Timezone:  tt.fields.Timezone,
				Alerts:    tt.fields.Alerts,
				Key:       tt.fields.Key,
				Excludes:  tt.fields.Excludes,
				client:    tt.fields.client,
			}
			if err := o.OneCallUsingCoordinates(tt.args.lat, tt.args.lon); (err != nil) != tt.wantErr {
				t.Errorf("OneCallUsingCoordinates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
