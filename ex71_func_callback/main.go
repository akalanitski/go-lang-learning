package main

import (
	"fmt"
	"math"
)

func printArea(area func(c float64) float64, c float64) {
	result := area(c)
	fmt.Println("Area is", result)
}

func areaOfSquare(s float64) float64 {
	return s * s
}

func areaOfCircle(c float64) float64 {
	return math.Pi * c * c
}

func main() {
	fmt.Println("Area of a square with side 10: ")
	printArea(areaOfSquare, 10)

	fmt.Println("Area of a circle with radius 10: ")
	printArea(areaOfCircle, 10)
}
