package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to get request")
	PerformgetRequest()
}

func PerformPostjsonRequest() {
	const myurl = ""

	requestBody := strings.NewReader(`
		{
			"coursename":"Learn Go"
			"price":"100"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformgetRequest() {
	const myurl = "https://www.google.com"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Response code : ", response.StatusCode)
	content, _ := ioutil.ReadAll((response.Body))

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = ""

	data := url.Values{}

	data.Add("firstName", "Ajay")
	data.Add("lastName", "Jeena")

	response, _ := http.PostForm(myurl, data)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
