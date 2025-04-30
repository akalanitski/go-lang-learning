package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Printf("Hello, world!\n")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		panic(1)
	}

	fmt.Printf("%d\n", n)
}
