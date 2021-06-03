/*

This file uses the NOAA API to return weather for a particular location.
It takes in latitude and longitude of a particular location.

Three API calls are made to the NOAA API.
(1) Returns the links that can be used to get weekly and hourly forecast
(2) API call is made to the weekly forecast link from the first API
(3) API call is made to the hourly forecast link from the first API

The hourly forecast and the weekly forecast are merged and returned back to the client.

*/

package weather

import (
	"encoding/json"
	"fmt"
	"github.com/nimitpatel26/tenki/go/geodata"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Following structs are used to store the metadata response (first call)
// from the NOAA API
type metaProperties struct {
	DailyForecast       string `json:"forecast"`
	ForecastHourly      string `json:"forecastHourly"`
	ForecastGridData    string `json:"forecastGridData"`
	ObservationStations string `json:"observationStations"`
	ForecastZone        string `json:"forecastZone"`
	County              string `json:"county"`
	FireWeatherZone     string `json:"fireWeatherZone"`
	TimeZone            string `json:"timeZone"`
	RadarStation        string `json:"radarStation"`
}

type weatherMetaResp struct {
	Properties metaProperties `json:"properties"`
}

// Following structs are used to store the response from the weekly and
// and the hourly forecast
type period struct {
	Name             string   `json:"name"`
	StartTime        string   `json:"startTime"`
	EndTime          string   `json:"endTime"`
	IsDaytime        bool     `json:"isDaytime"`
	Temperature      float64  `json:"temperature"`
	TemperatureUnit  string   `json:"temperatureUnit"`
	WindSpeed        string   `json:"windSpeed"`
	WindDirection    string   `json:"windDirection"`
	Icon             string   `json:"icon"`
	ShortForecast    string   `json:"shortForecast"`
	DetailedForecast string   `json:"detailedForecast"`
	Children         []period `json:"children"`
}

type elevation struct {
	Value    float64 `json:"value"`
	UnitCode string  `json:"unitCode"`
}

type properties struct {
	Updated           string    `json:"updated"`
	Units             string    `json:"units"`
	ForecastGenerator string    `json:"forecastGenerator"`
	GeneratedAt       string    `json:"generatedAt"`
	UpdateTime        string    `json:"updateTime"`
	ValidTimes        string    `json:"validTimes"`
	Elevation         elevation `json:"elevation"`
	Periods           []period  `json:"periods"`
}

type weatherResp struct {
	Properties properties `json:"properties"`
}

// Make the first API call to the NOAA API and
// return the response
func getWeatherMetaResp(lat float64, lon float64) *weatherMetaResp {
	target := fmt.Sprintf(os.Getenv("NOAA_API"), lon, lat)
	resp, err := http.Get(target)

	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	respBody := new(weatherMetaResp)
	err = json.NewDecoder(resp.Body).Decode(respBody)

	if err != nil {
		log.Fatal(err)
	}

	return respBody
}

// Make the second or third API call to get forecast data
// API (string) is used to specify which forecast call it is
func getWeatherResp(api string) *weatherResp {
	resp, err := http.Get(api)

	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	respBody := new(weatherResp)
	err = json.NewDecoder(resp.Body).Decode(respBody)

	if err != nil {
		log.Fatal(err)
	}

	return respBody
}

// Gets the weekly forecast and the hourly forecast and
// merges the result
func GetWeather(location string) properties {
	coordinates := geodata.GetLatLon(location)
	metadata := getWeatherMetaResp(coordinates[0], coordinates[1])
	weekly := getWeatherResp(metadata.Properties.DailyForecast)
	hourly := getWeatherResp(metadata.Properties.ForecastHourly)

	weekIndex := 0
	hourIndex := 0

	for weekIndex < len(weekly.Properties.Periods) {
		day := weekly.Properties.Periods[weekIndex]
		tDayStart, _ := time.Parse(time.RFC3339, day.StartTime)
		tDayEnd, _ := time.Parse(time.RFC3339, day.EndTime)

		for hourIndex < len(hourly.Properties.Periods) {
			hour := hourly.Properties.Periods[hourIndex]
			tHourStart, _ := time.Parse(time.RFC3339, hour.StartTime)
			tHourEnd, _ := time.Parse(time.RFC3339, hour.EndTime)

			if (tHourStart.Equal(tDayStart) || tHourStart.After(tDayStart)) && (tHourEnd.Equal(tDayEnd) || tHourEnd.Before(tDayEnd)) {
				weekly.Properties.Periods[weekIndex].Children = append(weekly.Properties.Periods[weekIndex].Children, hour)

			} else if tHourEnd.After(tDayStart) {

				break
			}

			hourIndex++
		}
		weekIndex++
	}
	return weekly.Properties
}
