package main

import (
	"fmt"
)

type engine struct {
	isElectric bool
}

type vehicle struct {
	engine
	make  string
	model string
	door  int
	color string
}

// T ake the code from the previous exercise, then store the VALUES of type person in a map with
// the KEY of last name. Access each value in the map. Print out the values, ranging over the
// slic
func main() {
	cars := []vehicle{
		{
			engine: engine{
				isElectric: true,
			},
			make:  "toyota",
			model: "RAV4",
			door:  5,
			color: "white",
		},
		{
			engine: engine{
				isElectric: false,
			},
			make:  "Chery",
			model: "Beat",
			door:  5,
			color: "yello",
		},
	}

	for _, car := range cars {
		fmt.Print("CAR: \"", car.make, " ", car.model, "\" is ")
		if car.isElectric {
			fmt.Println("electric")
		} else {
			fmt.Println("NOT electric")
		}
	}
}
