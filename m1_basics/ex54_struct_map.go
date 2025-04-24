package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

// T ake the code from the previous exercise, then store the VALUES of type person in a map with
// the KEY of last name. Access each value in the map. Print out the values, ranging over the
// slic
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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v)
		for i, v1 := range v.iceCream {
			fmt.Println(i, ":", v1)
		}
	}
}
