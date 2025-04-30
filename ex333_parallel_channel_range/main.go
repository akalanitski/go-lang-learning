package main

import (
	"fmt"
)

func sendToChanel(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("send %d to channel\n", i)
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)
	go sendToChanel(c)
	for v := range c { // subscribe on the channel updates
		fmt.Printf("receive \"%d\" from the chennel\n", v)
	}

}
