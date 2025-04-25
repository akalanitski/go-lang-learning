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

type person struct {
	firstName string
	age       int
}

func (p person) greating() {
	fmt.Println("Welcome", p.firstName)
}

func main() {
	p1 := person{
		firstName: "John",
		age:       42,
	}
	p1.greating()
}
