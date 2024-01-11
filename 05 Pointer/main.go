package main

import "fmt"

func main() {
	myVar := 50
	var ptr = &myVar
	fmt.Println(*ptr)
	fmt.Println(&ptr)
	fmt.Println(*(&ptr))
	fmt.Println(&(*(&myVar)))
	fmt.Println(&(*(&(*(&myVar)))))
}
