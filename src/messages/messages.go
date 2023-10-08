package messages

import "fmt"

var (
    WelcomeMsg = "\nWeather CLI v1.0.0."
    CommandsMsg = "Enter one of these commands:\n" +
        fmt.Sprintf("\t%-32s%s", "search", "Search for cities.\n") +
        fmt.Sprintf("\t%-32s%s", "daily <location>", "Daily forecast for the next week. Provide city name in place of <location>.\n") +
        fmt.Sprintf("\t%-32s%s", "hourly <location>", "Hourly forecast for the next 24 hours. Provide city name in place of <location>.\n") +
        fmt.Sprintf("\t%-32s%s", "exit", "Exit program.\n") 
)
