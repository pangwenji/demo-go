package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division")
	}
	return a / b, nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Starting program")
	panic("Something went wrong!") // 引发 panic
	fmt.Println("This line will not be executed")
}
