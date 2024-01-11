package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	val := 10
	for val > 0 {
		fmt.Println(val)
		val--
	}

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	for idx := 0; idx < len(days); idx++ {

		if days[idx] == "Friday" {
			break
		}
		if days[idx] == "Wednesday" {
			continue
		}
		fmt.Println(days[idx])
	}

	for i := range days {
		fmt.Println(days[i])
	}
	
}
