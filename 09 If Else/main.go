package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your number")

	number, _ := reader.ReadString('\n')

	if number == "10" {
		fmt.Println("Your number is 10")
	} else {
		fmt.Println("Your number is not 10")
	}
}
