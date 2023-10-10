package utils

import (
	"strings"
	"weather/src/cmd"
)

type Error struct {
    Message string
}

func ParseCmdArgs(input string) (command string, location string, err *Error) {
    parts := strings.Split(input, " ")

    if len(parts) < 2 {
        return "", "", &Error{Message: "ERROR: location parameter not entered."}
    }
    
    if !cmd.IsValidCmd(parts[0]) {
        return "", "", &Error{Message: "ERROR: input does not match any " +
            "of the existing commands."}
    }

    return parts[0], parts[1], nil
}
