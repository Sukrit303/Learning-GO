package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files")

	content := "LOREM mlsdmn knsdjln kk jklns jklv jkl ijklf sjk  jkl "

	file, err := os.Create("./file.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println(length)
	defer file.Close()
	readFile("./file.txt")
}

func readFile(fileName string) {
	dataRead, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	fmt.Println(string(dataRead))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
