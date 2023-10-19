package modes

import (
	"fmt"
	httpclient "weather/src/httpClient"
	"weather/src/utils"
)

type Hourly struct {
    Location string
}

func (h* Hourly) Handle() (error) {
    hours, err := httpclient.GetHourlyForecast(h.Location)    

    if err != nil {
        return err
    }

    for _, hour := range hours {
        fmt.Printf(
            "%-16s Condition: %-24s Temp. C: %-6v Precip. mm: %-6v Wind m/s: %-6.2f Wind dir.: %4s\n",
            hour.Time,
            hour.Condition.Text,
            hour.Temp_c,
            hour.Precip_mm,
            utils.ConvertKphToMps(hour.Wind_kph),
            hour.Wind_dir,
        )
    }

    return nil
}
