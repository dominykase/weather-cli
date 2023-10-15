package utils

import (
	"errors"
	"weather/src/cmd"
)

func ParseCmdArgs(args []string) (command string, location string, err error) {
    if len(args) == 0 {
        return cmd.Help, "", nil
    }

    if args[0] == cmd.Exit || args[0] == cmd.Help {
        return cmd.Exit, "", nil
    }

    if len(args) < 2 {
        return "", "", errors.New("ERROR: location not entered.") 
    }
    
    if !cmd.IsValidCmd(args[0]) {
        return "", "", errors.New("ERROR: input does not match any of the available commands.")
    }

    return args[0], args[1], nil
}
