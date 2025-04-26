package main

import "fmt"

func main() {
	a := 100
	p := &a
	fmt.Printf("Value of a: %d\n", a)
	fmt.Printf("Address: %p, refered value is: %d", p, *p)
}
