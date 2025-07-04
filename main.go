package main

import (
	"flag"
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

	location, err := util.ParseArgs()
	if err != nil {
		panic(fmt.Sprintf("Error parsing arguments: %v", err))
	}

	var days int
	flag.IntVar(&days, "days", 7, "number of days to get the weather forecast for")
	flag.Parse()

	weather, err := lib.MakeWeatherRequest(location, days, api_key)
	if err != nil {
		panic(fmt.Sprintf(util.PANIC_API_REQUEST_ERROR, location, err))
	}

	fmt.Println("Weather data:", weather)
}
