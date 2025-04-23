package main

import "fmt"

// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
// slice:
// "James" , "Miss" , "Bond" ,
// "Moneypenny" , "Shaken, not stirred" "I'm 008."
// Range over the records, then range over the data in each record.
func main() {
	states := [][]string{
		{`James`, `Miss`, `Bond`},
		{` Moneypenny`, ` Shaken, not stirred`, `I'm 008.`},
	}
	for i, v := range states {
		fmt.Printf("%d, %s\n", i, v)
	}
}
