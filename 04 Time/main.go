package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006"))

	createDate := time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC)

	fmt.Println(createDate)
}
