package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PublicIP struct {
	IP string `json:"ip"`
}

// Using the public IP service to get the public IP address, this will fetch the users public IP address
func GetPublicUserIp() (string, error) {
	resp, err := http.Get(IP_SVC_URL)
	if err != nil {
		return "", fmt.Errorf("error fetching fallback public IP: %w", err)
	}
	defer resp.Body.Close()

	var result PublicIP
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response body: %w", err)
	}
	return fmt.Sprintf("public ip: %s\n", result.IP), nil
}
