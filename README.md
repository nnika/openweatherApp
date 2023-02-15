

This is a simple app that extract the weather information from openweathermap.org App
To run this app you need to call `source setting\setenv.sh` to set the key as at `OWM_APP_KEY`
environment variable. This app access that key using `os.Getenv("OWM_APP_KEY")`

To use this program please run `go run main.go` then bring up the browser and type:
`http://localhost:8000/coordination?lat={latitude}&lon={longitude}` 

for example: `http://localhost:8000/coordination?lat=40.696011&lon=-73.993286` that will display this information:
 