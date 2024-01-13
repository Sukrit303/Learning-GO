package main

import "fmt"

func main() {
	greet()
	Adder(1,2,3,4,5,6,7,8,9)
	fmt.Println(add(1, 2))
	fmt.Println(sub(1, 2))
	fmt.Println(mul(1, 2))
	fmt.Println(div(1, 2))

}
func add(x int, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func mul(x, y int) int {
	return x * y
}
func div(x, y int) int {
	return x / y
}

func greet() {
	fmt.Println("Hello")
}

func Adder( values ...int)int{
	total :=0
	for _, val:range values {
		total += val
	}
	return total
}
