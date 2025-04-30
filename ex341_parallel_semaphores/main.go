package main

import "fmt"

func send10(c chan<- int, done chan bool) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	done <- true
}

func closeChannel(c chan<- int, done chan bool) {
	<-done
	<-done
	close(c)
}

func main() {
	c := make(chan int)
	done := make(chan bool)

	go send10(c, done)
	go send10(c, done)
	go closeChannel(c, done)

	for i := range c {
		fmt.Printf("%d was received from channel\n", i)
	}

}
