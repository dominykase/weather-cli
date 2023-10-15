package main

import (
	"fmt"
	"os"
	"weather/src/utils"
)

func handleError(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func main() {
    args := os.Args[1:]
   
    command, location, err := utils.ParseCmdArgs(args);

    handleError(err)

    mode, err := utils.CreateMode(command, location)
    err = mode.Handle()

    handleError(err)
}
