package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	ajay := User{"Ajay", "ajay@go.dev", true, 16}

	fmt.Println(ajay)
	fmt.Printf("Name is %+v", ajay)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
