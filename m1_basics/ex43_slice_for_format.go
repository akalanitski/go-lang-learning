package main

import "fmt"

// Using a COMPOSITE LITERAL:
// ● create a SLICE of TYPE int
// ● assign these 10 VALUES
//
//	42, 43, 44, 45, 46, 47, 48, 49, 50, 51
//
// ● Range over the slice and print the values out
func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v \t %T %v\n", s[i], s[i], i)
	}
}
