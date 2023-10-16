package utils

import (
	"errors"
	"weather/src/cmd"
	"weather/src/modes"
)

func CreateMode(mode string, location string) (modes.Mode, error) {
    switch mode {
    case cmd.Help:
        return &modes.Help{}, nil
    case cmd.Search:
        return &modes.Search{Location: location}, nil
    case cmd.Daily:
        return &modes.Daily{Location: location}, nil
    default:
        return nil, errors.New("Invalid mode")
    }
}
