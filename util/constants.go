package util

const (
	PANIC_ENV_VAR_NOT_SET    = "WEATHER_API_KEY environment variable is not set. set it in the .env file or your environment."
	PANIC_API_REQUEST_FAILED = "failed to fetch weather data: status code %d"
	PANIC_API_REQUEST_ERROR  = "error fetching weather forecast for %s: %v"
	PANIC_API_FAILED_TO_READ = "failed to read response body: %v"
)

const (
	BASE_URL   = "http://api.weatherapi.com/v1/forecast.json"
	IP_SVC_URL = "https://api.ipify.org?format=json"
)
