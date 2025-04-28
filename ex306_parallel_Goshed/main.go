package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func init() {

}

func main() {

	fmt.Printf("OS: %v, ARCH %v, CPU %v, goroutines %v, GOMAXPROCS: %v\n",
		runtime.GOOS, runtime.GOARCH, runtime.NumCPU(), runtime.NumGoroutine(), runtime.GOMAXPROCS(-1))

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("func(),\t number of goroutines", runtime.NumGoroutine())
			runtime.Gosched() // Yield to the main goroutine
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Main,\t number of goroutines", runtime.NumGoroutine())
		runtime.Gosched() // Yield to the goroutine
	}

	wg.Wait()
	fmt.Printf("End with CPU %v, goroutines %v\n", runtime.NumCPU(), runtime.NumGoroutine())
}
