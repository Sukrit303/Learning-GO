package main

import (
	"encoding/json"
	"fmt"
)

// private structure always defined by small letter
type message struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Content string `json:"content"`
	Subject string `json:"subject"`
	Time    string `json:"time,omitempty"`
}

func main() {
	fmt.Println("JSON CONSUME")
	ecnodingJSON()
	decodingJSON()
}
func ecnodingJSON() {
	messageBody := []message{
		{
			To:      "JOHN",
			From:    "Atlas",
			Content: "Hello",
			Subject: "Hi",
			Time:    "12:00",
		},
		{
			To:      "JANE",
			From:    "Aloo",
			Content: "Hello Bro",
			Subject: "Hi Bro",
		},
	}

	finalMessages, err := json.MarshalIndent(messageBody, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalMessages)
}

func decodingJSON() {
	jsonData := []byte(`
	{
		"to": "JOHN",
		"from": "Atlas",
		"content": "Hello",
		"subject": "Hi",        
		"time": "12:00"
	}
	`)

	var message message

	isValid := json.Valid(jsonData)
	if isValid {
		fmt.Println("JSON is valid")

		json.Unmarshal(jsonData, &message)

		fmt.Println(message)
	} else {
		fmt.Println("JSON is not valid")
	}

	var myDataInMap map[string]interface{}
	json.Unmarshal(jsonData, &myDataInMap)
	fmt.Printf("%s\n", myDataInMap)

	for key, value := range myDataInMap{
		fmt.Printf("key is %s and value is %v\n", key, value)
	}
	
}
