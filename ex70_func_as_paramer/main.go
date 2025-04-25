package main

import "fmt"

func outer() func() {
	return func() {
		fmt.Println("Hello World")
	}
}

func main() {
	f1 := outer()
	f1()
}
