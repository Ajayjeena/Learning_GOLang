package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name  string   `json:"coursename"`
	Price int      `json:"-"`
	Tags  []string `json:"tags"`
}

func main() {

	fmt.Println("Welcome to Json")

	encodeJson()
}

func encodeJson() {

	courses := []course{
		{"ReactJS Bootcamp", 299, []string{"Web-Dev", "JS"}},
		{"Angular Bootcamp", 299, []string{"Web-Dev", "JS"}},
	}

	finalJson, _ := json.MarshalIndent(courses, "", "\t")

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
	{
		"coursename": "ReactJS Bootcamp",   
		"tags": [
				"Web-Dev",
				"JS"
		]
	}
	`)

	var mycourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonData, &mycourse)

		fmt.Printf("%#v\n", mycourse)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

}
