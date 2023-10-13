package httpclient

import (
	"encoding/json"
	"errors"
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
        return nil, err
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("ERROR: API returned status code != 200. Exiting.") 
    }

    var cities []City

    if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
        return nil, err
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
