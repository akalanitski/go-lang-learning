package main

import "fmt"

func main() {
	c := make(chan int)
	a := <-c
	fmt.Printf("%d was received from channel\n", a)

	//fatal error: all goroutines are asleep - deadlock!
	//
	//goroutine 1 [chan send]:
	//main.main()
}
