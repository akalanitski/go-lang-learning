package main

import (
	"fmt"
	"sync"
)

var c = make(chan int, 2)
var wg sync.WaitGroup

func increment() {
	v := <-c
	v++
	c <- v
	fmt.Println("counter: ", v)
	wg.Done()
}

func main() {
	wg.Add(100)
	c <- 0
	for i := 0; i < 100; i++ {
		go increment()
	}
	wg.Wait()
	fmt.Printf("%v\n", <-c)
}
