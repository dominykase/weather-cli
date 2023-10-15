package main

import (
	"fmt"
	"os"
	"weather/src/cmd"
	httpclient "weather/src/httpClient"
	"weather/src/messages"
	"weather/src/utils"
)

func main() {
    args := os.Args[1:]
   
    command, location, err := utils.ParseCmdArgs(args);

    if err != nil {
        fmt.Println(err);
        return
    }

    switch (command) {
        case cmd.Exit:
            fmt.Println(messages.ExitMsg)
            break
        case cmd.Help:
            fmt.Printf(messages.CommandsMsg)
            break
        case cmd.Search:
            cities, err := httpclient.GetCities(location)
            
            if err != nil {
                panic(err)
            }

            fmt.Println("Found cities:")
            for _, city := range cities {
                fmt.Printf("\t%s (%s)\n", city.Name, city.Url)
            }
            fmt.Printf("\n")
            break;
        case cmd.Daily:
            break;
        case cmd.Hourly:
            break;
        default:
            break;
    }
}
