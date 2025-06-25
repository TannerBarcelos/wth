package util

import (
	"fmt"
	"os"
	"strings"
)

// ParseArgs checks if a location is provided as a command line argument. If not, it attempts to get the user's IP address to determine the location.
func ParseArgs() (string, error) {
	var location string
	if len(os.Args) > 1 {
		location = os.Args[1]
	} else {
		fmt.Println("No location provided. Using user's IP address to determine location.")
		ip, err := GetPublicUserIp()
		if err != nil {
			panic(fmt.Sprintf("Failed to get user IP: %v", err))
		}
		location = ip
	}
	if strings.TrimSpace(location) == "" {
		return "", fmt.Errorf("location cannot be empty")
	}
	return location, nil
}
