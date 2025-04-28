package main

import (
	"fmt"
)

var counter int

func increment() {
	counter = counter + 1
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}

	fmt.Scanln() // Wait user press any key
	fmt.Println("Counter:", counter)
}
