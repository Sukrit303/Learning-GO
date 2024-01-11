package main

import "fmt"

func main() {
	maps := make(map[string]int)
	maps["First"] = 1
	maps["Second"] = 2
	maps["Third"] = 3

	fmt.Println(maps)
}
