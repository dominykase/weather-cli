package httpclient

type City struct {
    Id int
    Name string
    Region string
    Country string
    Lat float64
    Lon float64
    Url string
}

type Day struct {
    MaxTemp_c float64
    MaxTemp_f float64
    MinTemp_c float64
    MinTemp_f float64
    AvgTemp_c float64
    AvgTemp_f float64
    MaxWind_mph float64
    MaxWind_kph float64
    TotalPrecip_mm float64
    TotalPrecip_in float64
    TotalSnow_cm float64
    AvgVis_km float64
    AvgVis_miles float64
    AvgHumidity float64
    DailyWillItRain int
    DailyChanceOfRain string
    DailyWillItSnow int
    DailyChanceOfSnow string
    Condition struct { 
        Text string
        Icon string
        Code int
    }
    Uv float64
}

type DailyForecastResponse struct {
    Forecast ForecastData `json:"forecast"`
}

type ForecastData struct {
    Days []ForecastDay `json:"forecastday"`
}

type ForecastDay struct {
    Date string 
    DateEpoch int
    Day Day
}

