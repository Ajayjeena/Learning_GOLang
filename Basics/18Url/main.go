package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://google.com?name='AJAY'"

func main() {
	fmt.Println("Welcome to URL handling")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams)
}
