package httpclient

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

const BaseUrl = "http://api.weatherapi.com"

type City struct {
    Id int
    Name string
    Region string
    Country string
    Lat float64
    Lon float64
    Url string
}

func GetCities(cityInput string) ([]City, error) {
    apiKey := getApiKey()
    resp, err := http.Get(BaseUrl + "/v1/search.json?key=" + apiKey + "&q=" + escapeNewLine(cityInput))

    if err != nil {
        panic("ERROR: something went wrong when fetching cities. Exiting.")
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        panic("ERROR: API returned wrong response code. Exiting.")
    }

    var cities []City

    if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
        panic("ERROR: failed to decode cities response body. Exiting.")
    }

    return cities, nil
}

func getApiKey() string {
    apiKey := os.Getenv("WEATHER_API_KEY")

	if apiKey == "" {
		panic("WEATHER_API_KEY environment variable is not set.")
    }

    return apiKey;
}

func escapeNewLine(input string) string {
    return strings.ReplaceAll(input, "\n", "")
}
