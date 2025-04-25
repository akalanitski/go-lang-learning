/*
“defer” multiple functions in main
○ show that a deferred func runs after the func containing it exits.
○ determine the order in which the multiple defer funcs run
*/

package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello")
}
func sayWorld() {
	fmt.Println("World")
}
func main() {
	defer sayHello()
	sayWorld()
}
