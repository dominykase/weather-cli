package messages

import "fmt"

var (
    CommandsMsg = "Weather CLI v1.0.0.\nAdd one of these arguments:\n" +
        fmt.Sprintf("\t%-32s%s", "search <location>", "Search for cities. Provide city name in place of <location>.\n") +
        fmt.Sprintf("\t%-32s%s", "daily <location>", "Daily forecast for the next week. Provide city name in place of <location>.\n") +
        fmt.Sprintf("\t%-32s%s", "hourly <location>", "Hourly forecast for the next 24 hours. Provide city name in place of <location>.\n") +
        fmt.Sprintf("\t%-32s%s", "exit", "Exit program.\n")
)
