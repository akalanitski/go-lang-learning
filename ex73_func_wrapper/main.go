package main

import (
	"fmt"
	"time"
)

func WrapperFunc(fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Println("Time elapsed: ", elapsed)
}

func MyFunc() {
	time.Sleep(2 * time.Second)
	fmt.Println("MyFunc")
}

func main() {
	WrapperFunc(MyFunc)
}
