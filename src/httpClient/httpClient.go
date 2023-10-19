package httpclient

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
)

const BaseUrl = "http://api.weatherapi.com"

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

func GetDailyForecast(cityInput string) ([]ForecastDay, error) {
    apiKey := getApiKey()
    resp, err := http.Get(BaseUrl + "/v1/forecast.json?key=" + apiKey + "&q=" + escapeNewLine(cityInput) + "&days=5&aqi=no&alerts=no")

    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("ERROR: API returned status code != 200. Exiting.") 
    }

    var data ForecastResponse

    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return nil, err
    }

    return data.Forecast.Days, nil
}

func GetHourlyForecast(cityInput string) ([]ForecastHour, error) {
    apiKey := getApiKey()
    resp, err := http.Get(BaseUrl + "/v1/forecast.json?key=" + apiKey + "&q=" + escapeNewLine(cityInput) + "&days=3&aqi=no&alerts=no")

    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("ERROR: API returned status code != 200. Exiting.") 
    }

    var data ForecastResponse 

    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return nil, err
    }

    var hours []ForecastHour

    for _, day := range data.Forecast.Days {
        hours = append(hours, day.Hours...)
    }

    return hours, nil
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
