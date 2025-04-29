// Sample of console Input / Outupt
package main

import "fmt"

func main() {
	var userName string                 // the place for name inside the program
	fmt.Println("Enter your name:")     // display promt for user
	fmt.Scanln(&userName)               // waiting for input
	fmt.Printf("Hello %s!\n", userName) // display greating
}
