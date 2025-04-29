package main

import (
	"fmt"
)

func sendToChanel(c chan<- int) {
	c <- 42
}

func receiveFromChanel(c <-chan int) {
	fmt.Println(<-c)
}

func main() {
	c := make(chan int, 2)
	sendToChanel(c)      // pass bi-directional chanel to reciever chanel
	receiveFromChanel(c) // pass bi-directional chanel to sender chanel
}
