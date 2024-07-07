package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files")

	content := "This needs to go in a file"

	file, err := os.Create("./myFile")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("length is ", length)
	defer file.Close()
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is ", string(databyte))
}
