package main

import "fmt"

func main() {
	var a int = 0
	var b int = 0
	var sum int = 0

	fmt.Println("Enter the first number:")
	fmt.Scan(&a)

	fmt.Println("Enter the second number:")
	fmt.Scan(&b)

	sum = a + b
	fmt.Printf("Sum is %d\n", sum)
}
