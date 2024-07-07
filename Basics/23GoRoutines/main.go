package main

import (
	"fmt"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup

func main() {

	go greeter("hello")
	wg.Add(1)
	go greeter("World")
	wg.Add(1)
	wg.Wait()
}

func greeter(s string) {
	defer wg.Done()
	for i := 1; i < 6; i++ {
		fmt.Println(s)
	}
}
