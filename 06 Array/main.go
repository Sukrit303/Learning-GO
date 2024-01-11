package main

import "fmt"

func main() {
	var data = []string{"10", "20", "30", "40"}
	

	fmt.Println(data)

	data = append(data[:2])
	fmt.Println(data)
}
