package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 42
	fmt.Printf("42 was sent to channel\n")

	//fatal error: all goroutines are asleep - deadlock!
	//
	//goroutine 1 [chan send]:
	//main.main()
}
