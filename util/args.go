package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CliArgs struct {
	Location string
	Days     string
}

// ParseArgs checks if a location is provided as a command line argument. If not, it attempts to get the user's IP address to determine the location.
func ParseArgs() (*CliArgs, error) {
	var location string
	var days = "7" // Default to 7 days if not specified
	if len(os.Args) > 1 {
		location = os.Args[1]
		// user provided a location and days to forecast
		if len(os.Args) > 2 {
			days = os.Args[2]
			coercedDays, err := strconv.Atoi(days)
			if err != nil {
				return nil, fmt.Errorf("invalid number of days: %v", err)
			}
			if coercedDays < 1 || coercedDays > 14 {
				return nil, fmt.Errorf("number of days must be between 1 and 14")
			}
			days = strconv.Itoa(coercedDays)
		}
	} else {
		fmt.Println("No location provided. Using user's IP address to determine location.")
		ip, err := GetPublicUserIp()
		if err != nil {
			panic(fmt.Sprintf("Failed to get user IP: %v", err))
		}
		location = ip
	}
	if strings.TrimSpace(location) == "" {
		return nil, fmt.Errorf("location cannot be empty")
	}
	return &CliArgs{
		Location: strings.TrimSpace(location),
		Days:     days,
	}, nil
}
