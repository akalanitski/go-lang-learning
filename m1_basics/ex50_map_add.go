package main

import "fmt"

// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
// slice:
// "James" , "Miss" , "Bond" ,
// "Moneypenny" , "Shaken, not stirred" "I'm 008."
// Range over the records, then range over the data in each record.
func main() {
	m := map[string][]string{
		"bond_james":       {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}
	for k, v := range m {
		fmt.Printf("%s, %s\n", k, v)
	}
	m["fleming_ian"] = []string{"steaks", "cigars", "espionage"}
	fmt.Println("------------")
	for k, v := range m {
		fmt.Printf("%s, %s\n", k, v)
	}
}
