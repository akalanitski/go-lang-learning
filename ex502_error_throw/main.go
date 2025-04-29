package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	n, err := sqrt(-10.23)
	if err != nil {
		log.Println("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", n)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math error: negative number")
	}
	return 42, nil
}
