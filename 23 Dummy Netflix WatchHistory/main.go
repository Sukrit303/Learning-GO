package main

import (
	"BackendNetflixWatchHistory/routes"
	"fmt"
	"log"
	"net/http"
)



func main() {
	fmt.Println("Server is starting")
	r := routes.Router()
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Server is Running")
}
