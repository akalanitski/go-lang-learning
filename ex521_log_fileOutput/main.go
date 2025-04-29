package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile, err := os.Create("myLog.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	myFile, err := os.Open("NoFile.txt")
	if err != nil {
		log.Println("Error Happend", err)
	}
	defer myFile.Close()
	fmt.Printf("Hello world\n")
}
