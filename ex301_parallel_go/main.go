package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func heavyTask(id int) {
	start := time.Now()
	sum := 0
	for i := 0; i < 4_000_000_000; i++ { // Heavy loop
		sum += i
	}
	fmt.Printf("Goroutine %d finished in %v\n", id, time.Since(start))
	wg.Done()
}

func main() {

	runtime.GOMAXPROCS(2) // Limit to 1 CPU core
	fmt.Printf("OS: %v, ARCH %v, CPU %v, goroutines %v, GOMAXPROCS: %v\n",
		runtime.GOOS, runtime.GOARCH, runtime.NumCPU(), runtime.NumGoroutine(), runtime.GOMAXPROCS(-1))

	wg.Add(3)
	start := time.Now()

	go heavyTask(1)
	go heavyTask(2)
	go heavyTask(3)

	// Wait enough time for both goroutines to finish

	wg.Wait()
	fmt.Printf("Total elapsed time: %v\n", time.Since(start))
}
