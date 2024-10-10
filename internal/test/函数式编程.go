package main

import "fmt"

type operation func(a, b int) int

func calulate(x, y int, op operation) int {
	return op(x, y)
}

func multip(a int) func(b int) int {
	return func(b int) int {
		return a * b
	}
}

func main() {
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("calulate", calulate(1, 2, add))
	multiply := func(a, b int) int {
		return a * b
	}

	fmt.Println("multiply", calulate(1, 2, multiply))
	triple := multip(3)
	fmt.Println(triple(5))
}
