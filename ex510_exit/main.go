package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello, world!\n")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", n)
	os.Exit(1)
}
