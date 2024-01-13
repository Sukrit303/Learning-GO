package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	GetApiCallReq()
}
func GetApiCallReq() {
	const myURL = "http://localhost:7777"
	res, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println(res.StatusCode)

	constent, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(constent))
}
