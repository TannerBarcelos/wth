package lib

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tannerbarcelos/wth/util"
)

// Given a location and an API key, this function makes a request to the Openweather API. Returns the weather data as a string or an error if the request fails.
func MakeWeatherRequest(location, days, api_key string) (string, error) {
	encodedLocation := url.QueryEscape(location)
	url := fmt.Sprintf("%s?key=%s&q=%s&days=%s", util.BASE_URL, api_key, encodedLocation, days)

	res, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf(util.PANIC_API_REQUEST_FAILED, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf(util.PANIC_API_FAILED_TO_READ, err)
	}

	return string(body), nil
}
