package utils

import (
	"errors"
	"weather/src/cmd"
)

func ParseCmdArgs(args []string) (command string, location string, err error) {
    if len(args) == 0 {
        return cmd.Help, "", nil
    }

    if !cmd.IsValidCmd(args[0]) {
        return "", "", errors.New("ERROR: input does not match any of the available commands.")
    }

    if len(args) < 2 {
        return "", "", errors.New("ERROR: location not entered.") 
    }
    
    return args[0], args[1], nil
}
