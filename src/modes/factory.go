package modes 

import (
	"errors"
	"weather/src/cmd"
)

func CreateMode(mode string, location string) (Mode, error) {
    switch mode {
    case cmd.Help:
        return &Help{}, nil
    case cmd.Search:
        return &Search{Location: location}, nil
    case cmd.Daily:
        return &Daily{Location: location}, nil
    case cmd.Hourly:
        return &Hourly{Location: location}, nil
    default:
        return nil, errors.New("Invalid mode")
    }
}
