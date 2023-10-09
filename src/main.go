package main

import (
    "fmt"
    "weather/src/messages"
)

func main() {
    var input string;

    for input != "exit" {
        switch (input) {
            case "":
                fmt.Println(messages.WelcomeMsg);
                break;
            case "search": 
                break;
            case "daily":
                break;
            case "hourly":
                break;
            default:
                fmt.Println(
                    fmt.Sprintf(
                        "\n\"%s\" is not a recognised input.\n\n",
                        input,
                    ),
                );
                break;
        }

        fmt.Println(messages.CommandsMsg);
        fmt.Scanln(&input);
    }
}
