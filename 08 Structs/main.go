package main

import "fmt"

func main() {
	user1 := User{"John", "j@j.com",30}
	fmt.Println(user1)
}

type User struct {
	name  string
	email string
	age   int
}
