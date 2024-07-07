package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Tomato"}

	fruitList = append(fruitList, "Peach")

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 9
	highScores[3] = 8
	fmt.Println(highScores)
}
