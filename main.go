package main

import "net/http"

func main() {
	url := "http://example.com"
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic("Failed to fetch the URL: " + url + " with status code: " + res.Status)
	}
}
