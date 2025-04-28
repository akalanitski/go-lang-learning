package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func increment() {
	mu.Lock()
	counter = counter + 1
	mu.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}

	fmt.Scanln() // Wait user press any key
	fmt.Println("Counter:", counter)
}
