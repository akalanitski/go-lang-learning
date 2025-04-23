package main

import "fmt"

// Using the code from the previous example, use SLICING to create the following new slices
// which are then printed:
// ● [42 43 44 45 46]
// ● [47 48 49 50 51]
// ● [44 45 46 47 48]
// ● [43 44 45 46 47]
func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(s[0:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:7])
	fmt.Println(s[1:6])
}
