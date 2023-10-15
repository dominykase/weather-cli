package cmd

const (
    Exit string = "exit"
    Help string = "help"
    Search string = "search"
    Daily string = "daily"
    Hourly string = "hourly"
)

func IsValidCmd(input string) bool {
	validEnums := map[string]bool{
        Exit: true,
        Help: true,
        Search: true,
		Daily:  true,
		Hourly: true,
	}

	return validEnums[input]
}
