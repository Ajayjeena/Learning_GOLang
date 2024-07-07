package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointer lesson")

	myNumber := 3
	var ptr = &myNumber

	fmt.Println(ptr, *ptr)
}
