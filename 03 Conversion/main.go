package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hi enter number")
	number, err := reader.ReadString('\n')
	fmt.Println("The number is", number, "and error is", err)
	number1, err := strconv.ParseFloat(strings.TrimSpace(number), 64)
	fmt.Println("The number is", number1, "and error is", err)
}
