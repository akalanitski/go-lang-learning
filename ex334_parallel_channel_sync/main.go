package main

import (
	"fmt"
)

func sendToChanel(c chan<- int, ack <-chan string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("send %d to channel\n", i)
		c <- i
		<-ack // wait for receiver to acknowledge
	}
	close(c)
}

func main() {
	c := make(chan int)
	ack := make(chan string)
	go sendToChanel(c, ack)
	for v := range c { // subscribe on the channel updates
		fmt.Printf("receive \"%d\" from the chennel\n", v)
		ack <- "received"
	}

}
