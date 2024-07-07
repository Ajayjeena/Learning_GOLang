package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops classes")

	days := []string{"Monday", "Tuesday", "Wednesday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	count := 1

	for count < 10 {
		fmt.Println("The value is ", count)
		count++
	}
}
