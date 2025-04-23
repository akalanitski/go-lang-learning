package main

import "fmt"

// For this exercise, do the following:
// ● Create a slice to store the names of all of the states in the United States of America.
// ○ Use make and append to do this.
// ○ Goal: do not have the array that underlies the slice created more than once.
// ● Print out
// ○ the len
// ○ the cap
// ○ the values, along with their index position, without using the range clause.
// ● Here is a list of the 50 states:
func main() {
	states := []string{`Alabama`, `Alaska`, `Arizona`, ` Arkansas`, ` California`,
		` Colorado`, `Connecticut`, `Delaware`, ` Florida`, `Georgia`}
	fmt.Println(len(states), cap(states))
	for i, v := range states {
		fmt.Printf("%d, %s\n", i, v)
	}
}
