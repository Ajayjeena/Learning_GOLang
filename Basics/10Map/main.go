package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["RB"] = "Ruby"

	for key, value := range languages {
		fmt.Println("For key %v , the value is %v\n", key, value)
	}
	delete(languages, "RB")

	fmt.Println(languages)
}
