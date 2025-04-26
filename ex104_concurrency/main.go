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

	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("ARCH: ", runtime.GOARCH)
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Grourutine: ", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Grourutine: ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Grourutine: ", runtime.NumGoroutine())
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}
