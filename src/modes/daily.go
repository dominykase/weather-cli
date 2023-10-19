package modes

import (
	"fmt"
	httpclient "weather/src/httpClient"
	"weather/src/utils"
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
        fmt.Printf(
            "%-12s Max temp. C: %-6v Min temp. C: %-6v Avg. temp C: %-6v Max wind m/s: %-6.2f Precip mm: %-6v Snow cm: %-6v\n",
            forecastDay.Date,
            forecastDay.Day.MaxTemp_c,
            forecastDay.Day.MinTemp_c,
            forecastDay.Day.AvgTemp_c,
            utils.ConvertKphToMps(forecastDay.Day.MaxWind_kph),
            forecastDay.Day.TotalPrecip_mm,
            forecastDay.Day.TotalSnow_cm,
        )
    }

    return nil
}
