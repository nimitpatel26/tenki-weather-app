/*

This package gets the news data for a particular location
using the NewsAPI.

*/

package news

import (
	"encoding/json"
	"fmt"
	"github.com/nimitpatel26/tenki/go/geodata"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// The following structs are used to store
// the response from the NewsAPI
type source struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type article struct {
	Source      source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type newsResp struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []article `json:"articles"`
}

// Takes in place, district, and region and returns a formatted
// query string that can be used by the NewAPI
func getQuery(place string, dist string, region string) string {
	query := "\"" + place + "\" OR \"" + dist + "\""

	if len(place) == 0 && len(dist) == 0 {
		query = "\"" + region + "\""
	} else if len(place) == 0 {
		query = "\"" + dist + "\""
	} else if len(dist) == 0 {
		query = "\"" + place + "\""
	}
	return url.PathEscape(query)
}

// Calculates the date that earliest news article can have
// Currently, only articles that are a week old will be returned
func getLatestDate() string {
	today := time.Now().UTC()
	lastWeek := today.AddDate(0, 0, -7)
	return fmt.Sprintf("%v", lastWeek.Format("2006-01-02"))
}

// Takes in a location string and returns a NewAPI url
// that can be called to get the news data
func getNewsUrl(location string) string {
	place, dist, region := geodata.GetAddress(location)
	query := getQuery(place, dist, region)
	latestDate := getLatestDate()
	return fmt.Sprintf(os.Getenv("NEWS_API"), query, latestDate, os.Getenv("NEWS_API_KEY"))
}

// Calls the NewsAPI and places the result in the struct
func getNewsResp(newsUrl string) *newsResp {
	resp, err := http.Get(newsUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	respBody := new(newsResp)
	err = json.NewDecoder(resp.Body).Decode(respBody)

	if err != nil {
		log.Fatal(err)
	}

	return respBody
}

// Takes in a location string and returns
// a list of articles for that location
func GetNews(location string) []article {
	newsUrl := getNewsUrl(location)
	return getNewsResp(newsUrl).Articles
}
