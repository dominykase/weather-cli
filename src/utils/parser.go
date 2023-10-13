package utils

import (
	"errors"
	"strings"
	"weather/src/cmd"
)

func ParseCmdArgs(input string) (command string, location string, err error) {
    parts := strings.Split(input, " ")

    if len(parts) < 2 {
        return "", "", errors.New("ERROR: location not entered.") 
    }
    
    if !cmd.IsValidCmd(parts[0]) {
        return "", "", errors.New("ERROR: input does not match any of the available commands.")
    }

    return parts[0], parts[1], nil
}
