package main

import (
	"bufio"
	"fmt"
	"os"
	"weather/src/cmd"
	httpclient "weather/src/httpClient"
	"weather/src/messages"
	"weather/src/utils"
)

func main() {
    var input string;
    
    reader := bufio.NewReader(os.Stdin)

    fmt.Println(messages.WelcomeMsg)
   
    for input != "exit" {
        fmt.Println(messages.CommandsMsg)
        input, stdinErr := reader.ReadString('\n')
        
        if stdinErr != nil {
            panic("ERROR: something went wrong with input. Exiting.")
        }

        command, location, err := utils.ParseCmdArgs(input);

        if err != nil {
            fmt.Println((*err).Message);

            continue;
        }

        switch (command) {
            case cmd.Search:
                cities, err := httpclient.GetCities(location)
                
                if err != nil {
                    panic(err)
                }

                for _, city := range cities {
                    fmt.Printf("%s (%s)\n", city.Name, city.Url)
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
}
