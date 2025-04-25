package main

import "fmt"

func main() {
	a := 1
	f := func() {
		fmt.Println("Value of a: ", a)
		a++
	}
	for i := 0; i < 10; i++ {
		f()
	}
}
