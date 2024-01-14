package main

import (
	"encoding/json"
	"fmt"
)

type profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role string `json:"role"`
	Tags []string
}

func main() {
	fmt.Println("JSON")
	EncodingJSON()

}

func EncodingJSON() {
	userProfile := []profile{
		{
			Name: "John",
			Age:  30,
			Role: "Admin",
			Tags: []string{"WEB", "APP"},
		},
		{
			Name: "Jane",
			Age:  25,
			Role: "User",
			Tags: []string{"APP", "DEV"},
		},
	}

	// Packaging data into JSON
	finalJSON, err := json.MarshalIndent(userProfile, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJSON)
}
