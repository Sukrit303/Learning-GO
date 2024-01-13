package main

import "fmt"

// Defer keyword is used to defer the execution of a function until the surrounding function returns.
// LIFO order of execution of line function etc
func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("World")
	defer fmt.Println("!")
	fmt.Println("Namaste")
	defer fmt.Println("Bonjure")
}
