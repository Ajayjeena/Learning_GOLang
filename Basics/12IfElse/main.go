package main

import "fmt"

func main() {
	fmt.Println("if and else in golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "regular"
	} else {
		result = "Somrething else"
	}

	println(result)
}
