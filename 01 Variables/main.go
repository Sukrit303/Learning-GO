package main

import "fmt"

// Public varible
const AuthServToken = "123456789"

func main() {
	name := "John"
	class := 5
	istrue := true

	fmt.Println(name, class, istrue)
	var HI string = "Hello World"
	var HI1 bool = false
	var hi5 int = 5
	var hi6 float64 = 5.511111111111111111111111111111111
	fmt.Println(HI, HI1, hi5, hi6)

	fmt.Println(AuthServToken)
}
