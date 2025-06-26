# WTH

## Introduction
WTH is a CLI tool that allows a user to quickly get the weather forecast for a given location. It uses the weatherapi API to fetch the weather data and displays it in a user-friendly format.

## Installation
To build and run the cli, you need to have Go installed on your machine. You can download it from [here](https://golang.org/dl/).

To install WTH, run the following command:

```bash
go install github.com/tbarcelo/wth@latest
```

Next you you need to rename the `.env.example` file to `.env` and set your OpenWeatherMap API key in it. You can get a free API key by signing up at [OpenWeatherMap](https://app.swaggerhub.com/apis-docs/WeatherAPI.com/WeatherAPI/1.0.2).

## Usage
Once installed, you can use the `wth` command to get the weather forecast for a specific location. The command takes the location as an argument and returns the current weather conditions.

```bash
wth <location>
``` 

For example, to get the weather for New York City, you would run:

```bash
wth "New York City"
```

You can also specify the amount of days for the forecase by using the `-d` or `--days` flag:

```bash
wth "New York City" -d 5
```

> If you do not provide a location, WTH will use your current location based on your IP address (this requires an internet connection) and default to a 7 day forecast.

## Configuration
You can configure WTH to use a specific OpenWeatherMap API key by setting the `WTH_API_KEY` environment variable. If this variable is not set, WTH will use the default API key provided in the code.

## License
WTH is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.