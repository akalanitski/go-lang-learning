/*
* create a func with the identifier foo that
  - takes in a variadic parameter of type int
  - pass in a value of type []int into your func (unfurl the []int)
  - returns the sum of all values of type int passed in

create a func with the identifier bar that
  - takes in a parameter of type []int
  - returns the sum of all values of type int passed in+
*/
package ex62_func_interface_test

import (
	"fmt"
	"math"
	"strconv"
)

type square struct {
	width  float64
	height float64
}

func (s square) area() float64 {
	return s.width * s.height
}

func (s square) name() string {
	strWidth := strconv.FormatFloat(s.width, 'f', -1, 64)
	strHeight := strconv.FormatFloat(s.height, 'f', -1, 64)
	return "Square" + strWidth + "x" + strHeight + ": "
}

type circle struct {
	radius int
}

func (c circle) area() float64 {
	return math.Pi * float64(c.radius) * float64(c.radius)
}

func (c circle) name() string {
	strRadius := strconv.Itoa(c.radius)
	return "Circle" + strRadius + ": "
}

type shape interface {
	area() float64
	name() string
}

func printArea(s shape) {
	strArea := strconv.FormatFloat(s.area(), 'f', 2, 64)
	fmt.Println("Area of ", s.name(), "\t", strArea)
}

func main() {
	s := square{width: 10, height: 5}
	printArea(s)

	c := circle{radius: 5}
	printArea(c)
}
