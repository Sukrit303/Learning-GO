package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://loc.dev:3000/learn?yes"

func main() {
	fmt.Println(myURL)

	// Parsing in url

	result, err := url.Parse(myURL)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
}
