package modes

import (
	"fmt"
	httpclient "weather/src/httpClient"
)

type Search struct {
    Location string
}

func (m *Search) Handle() error { 
    cities, err := httpclient.GetCities(m.Location)
    
    if err != nil {
        return err
    }

    fmt.Println("Found cities:")
    for _, city := range cities {
        fmt.Printf("\t%s (%s)\n", city.Name, city.Url)
    }
    fmt.Printf("\n")
 
    return nil 
}
