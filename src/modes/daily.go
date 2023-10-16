package modes

import (
	"fmt"
	httpclient "weather/src/httpClient"
)

type Daily struct {
    Location string
}

func (d *Daily) Handle() (error) {
    fmt.Println("Daily forecast for " + d.Location + ":")
    forecasts, err := httpclient.GetDailyForecast(d.Location)

    if err != nil {
        return err
    }

    for _, forecastDay := range forecasts {
        fmt.Println(forecastDay.Date)
        fmt.Printf(
            "Max temp celsius: %v, Min temp celsius: %v, Avg temp celsius: %v\n",
            forecastDay.Day.MaxTemp_c,
            forecastDay.Day.MinTemp_c,
            forecastDay.Day.AvgTemp_c,
        )
        fmt.Printf(
            "Max wind kph: %v, Total precip mm: %v, Total snow cm: %v\n",
            forecastDay.Day.MaxWind_kph,
            forecastDay.Day.TotalPrecip_mm,
            forecastDay.Day.TotalSnow_cm,
        )
    }

    return nil
}
