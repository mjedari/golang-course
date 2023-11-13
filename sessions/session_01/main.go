package main

import (
	"fmt"
	"session_01/calculator"
)

func main() {
	number1 := 90
	number2 := 10

	result1 := calculator.Sum(number1, number2)
	result2 := calculator.Minus(number2, number1)

	fmt.Println("Result: ", result1, result2)

	number2 *= 2
	number2 = number2 * 2

	fmt.Println("Result: ", number2)

	// Comparison
	if number1 > number2 {
		fmt.Println("equal")
	}
}
