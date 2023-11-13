package main

import (
	"errors"
	"fmt"
	"log"
	"session_02/calculator"
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

	// zero values
	// int, flout, string, bool
	example := `test0
	test1
	test2
	`
	fmt.Printf("example is %v and you %v see", example, number2)
	log.Printf("example is %v and you %v see", example, number2)

	//number, err := fmt.Printf("example is %v and you %v see\n", example, number2)
	//fmt.Println(number, err)

	// adult checker
	// :=    =>    =
	// :=    =>   :=
	age := 19
	age, err := AdultCheckerV2(age)
	if err != nil {
		fmt.Println("the age is not valid")
		return
	}

	// blank identifier
	_, _ = fmt.Printf("the age: %v is valid", age)
}

func AdultCheckerV1(age int) bool {
	return age < 18
}

func AdultCheckerV2(age int) (int, error) {
	if age < 18 {
		return age, errors.New("age is under 18")
	}

	// null , nil
	return age, nil
}
