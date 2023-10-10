package main

import (
	"bufio"
	"fmt"
	"os"
	"weather/src/cmd"
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
                fmt.Println(location)
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
