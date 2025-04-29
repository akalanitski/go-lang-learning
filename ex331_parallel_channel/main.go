package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	r := make(<-chan int, 2) // recive chanel
	s := make(chan<- int, 2) // send chanel

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", r)
	fmt.Printf("%T\n", s)
}
