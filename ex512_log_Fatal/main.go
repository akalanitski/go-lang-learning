package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("myLogs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Hello world\n")
}
