package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/tannerbarcelos/wth/lib"
	"github.com/tannerbarcelos/wth/util"
)

func main() {
	api_key, ok := os.LookupEnv("WTH_API_KEY")
	if !ok {
		panic(util.PANIC_ENV_VAR_NOT_SET)
	}

	weatherArgs, err := util.ParseArgs()
	if err != nil {
		panic(fmt.Sprintf("Error parsing arguments: %v", err))
	}

	weather, err := lib.MakeWeatherRequest(weatherArgs.Location, weatherArgs.Days, api_key)
	if err != nil {
		panic(fmt.Sprintf(util.PANIC_API_REQUEST_ERROR, weatherArgs.Location, err))
	}

	fmt.Println("Weather data:", weather)
}
