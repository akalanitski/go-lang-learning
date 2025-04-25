package main

import (
	"fmt"
)

func foo() int {
	return 10
}

func bar() string {
	return "bar"
}

// create a func with the identifier foo that returns an int
// * create a func with the identifier bar that returns an int and a string
// * call both funcs
// * print out their results
func main() {
	fmt.Println("foo()", foo())
	fmt.Println("bar()", bar())
}
