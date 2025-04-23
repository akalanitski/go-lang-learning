package main

import "fmt"

// Follow these steps:
// ● start with this slice
// ○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// ● append to that slice this value
// ○ 52
// ● print out the slice
// ● in ONE STATEMENT append to that slice these values
// ○ 53
// ○ 54
// ○ 55
// ● print out the slice
// ● append to the slice this slice
// ○ y := []int{56, 57, 58, 59, 60}
func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	s = append(s, 52)
	fmt.Println(s)
	s = append(s, 53, 54, 55)
	fmt.Println(s)
	y := []int{56, 57, 58, 59, 60}
	s = append(s, y...)
	fmt.Println(s)

}
