package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64 = 0
var wg sync.WaitGroup

func increment() {
	atomic.AddInt64(&counter, 1)
	fmt.Println("counter: ", atomic.LoadInt64(&counter))
	wg.Done()
}

func main() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go increment()
	}
	wg.Wait()
	fmt.Println("In the End Counter:", counter)
}
