package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	const url = "https://www.google.com"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response from URL is type %T \n", response.Status)
	// Alaways close the connection
	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println(content)
}
