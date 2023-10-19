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

type ForecastResponse struct {
    Location LocationData `json:"location"`
    Forecast ForecastData `json:"forecast"`
}

type LocationData struct {
    Name string
    Region string
    Country string
    Lat float64
    Lon float64
    Timezone string `json:"tz_id"`
    Localtime string
}

type ForecastData struct {
    Days []ForecastDay `json:"forecastday"`
}

type ForecastDay struct {
    Date string 
    DateEpoch int
    Day Day
    Hours []ForecastHour `json:"hour"`
}

type ForecastHour struct {
    Time string
    Temp_c float64
    Condition struct {
        Text string
        Icon string
        Code int
    }
    Wind_kph float64
    Wind_dir string
    Precip_mm float64
}

