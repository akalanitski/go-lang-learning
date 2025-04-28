package main

import (
	"fmt"
	"math"
)

func addT[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Printf("Sum of 10 and 7 is %d\n", addT(10, 7))
	fmt.Printf("Sum of 11.5 and math.Pi is %v\n", addT(11.5, math.Pi))
}
