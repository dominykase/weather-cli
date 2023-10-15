package cmd

const (
    Help string = "help"
    Search string = "search"
    Daily string = "daily"
    Hourly string = "hourly"
)

func IsValidCmd(input string) bool {
	validEnums := map[string]bool{
        Help: true,
        Search: true,
		Daily:  true,
		Hourly: true,
	}

	return validEnums[input]
}
