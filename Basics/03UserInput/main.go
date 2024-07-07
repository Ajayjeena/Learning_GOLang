package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating =  ", input)
}
