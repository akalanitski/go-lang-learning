/*
Ex.5.4 Deitl
Finding the maximum of three integers
*/
package main

import "fmt"

func minimum(param ...int) int {
	if len(param) == 0 {
		return 0
	}
	result := param[0]
	for _, i := range param {
		if i < result {
			result = i
		}
	}
	return result
}

func main() {
	input := []int{0, 0, 0}
	fmt.Println("Enter 3 integers: ")
	fmt.Scanf("%d %d %d", &input[0], &input[1], &input[2])

	fmt.Println(minimum(input...))
}
