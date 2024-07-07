package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Learning Channel")

	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		//val,isChanelOpen := <-myCh
		fmt.Println(<-myCh)
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
