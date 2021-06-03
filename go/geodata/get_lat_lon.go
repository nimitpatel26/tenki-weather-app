/*

This file uses MapBox API to give latitude and logitude of
a particular location using the MapBox API

*/
package geodata

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Following structs are used to store response from the MapBox API
type context struct {
	Id   string `json:"Id"`
	Text string `json:"text"`
}

type feature struct {
	Text      string    `json:"text"`
	PlaceName string    `json:"place_name"`
	Center    []float64 `json:"center"`
	Context   []context `json:"context"`
}

type mapBoxResp struct {
	Type     string    `json:"type"`
	Features []feature `json:"features"`
}

// Gets the response back from MapBox
func getMBoxResp(location string) *mapBoxResp {
	target := fmt.Sprintf(os.Getenv("MAPBOX_API"), location, os.Getenv("MAPBOX_API_KEY"))
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

	respBody := new(mapBoxResp)
	err = json.NewDecoder(resp.Body).Decode(respBody)

	if err != nil {
		log.Fatal(err)
	}

	return respBody
}

// Extracts latitude and longitude from MapBox response
func getLatLon(resp *mapBoxResp) []float64 {
	locationFilter := "United States"

	for _, v := range resp.Features {

		targetFound := strings.Contains(v.PlaceName, locationFilter)

		if targetFound {
			return v.Center
		}

	}

	return nil
}

// For a given location, it returns the latitude and logitude using the MapBox API
func GetLatLon(location string) []float64 {
	resp := getMBoxResp(location)
	return getLatLon(resp)
}

// For a given location, it returns the place, district, and region
func GetAddress(location string) (string, string, string) {
	resp := getMBoxResp(location)
	place, dist, region, country := "", "", "", ""
	for _, feat := range resp.Features {
		for _, con := range feat.Context {

			if strings.Contains(con.Id, "place") {
				place = con.Text
			}

			if strings.Contains(con.Id, "district") {
				dist = con.Text
			}

			if strings.Contains(con.Id, "region") {
				region = con.Text
			}

			if strings.Contains(con.Id, "country") {
				country = con.Text
			}
		}
		if country == "United States" {
			return place, dist, region
		}

	}

	return "", "", ""

}
