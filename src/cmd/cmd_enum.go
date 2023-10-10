package cmd

const (
    Search string = "search"
    Daily string = "daily"
    Hourly string = "hourly"
)

func IsValidCmd(input string) bool {
	validEnums := map[string]bool{
		Search: true,
		Daily:  true,
		Hourly: true,
	}

	return validEnums[input]
}
