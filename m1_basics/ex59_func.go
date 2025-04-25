/*
* create a func with the identifier foo that
	* takes in a variadic parameter of type int
	* pass in a value of type []int into your func (unfurl the []int)
	* returns the sum of all values of type int passed in
create a func with the identifier bar that
	* takes in a parameter of type []int
	* returns the sum of all values of type int passed in+
*/

package main

import (
	"fmt"
)

func sum(params ...int) int {
	result := 0
	for _, param := range params {
		result += param
	}
	return result
}

func sumSlice(params []int) int {
	result := 0
	for _, param := range params {
		result += param
	}
	return result
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("sum(1, 2, 3, 4, 5)\n", sum(1, 2, 3, 4, 5))
	fmt.Println("sum(s...)\n", sum(s...))
	fmt.Println("sumSlice(s)\n", sumSlice(s))
}
