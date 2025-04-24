package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

// Create your own type “person” which will have an underlying type of “struct” so that it can store
// the following data:
// ● first name
// ● last name
// ● favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
// which stores the favorite flavors.
func main() {
	p1 := person{
		firstName: "John",
		lastName:  "Doe",
		iceCream:  []string{"Pistazien", "Erdberry"},
	}
	p2 := person{
		firstName: "Marry",
		lastName:  "Jein",
		iceCream:  []string{"Himberry", "Vannila"},
	}

	fmt.Println(p1, p2)
}
