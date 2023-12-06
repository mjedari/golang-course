package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Simple Calculator")

	for {
		fmt.Println("------------------")
		firstNumber := getFirstNumber()
		operator := getOperator()
		secondNumber, err := getSecondNumber()
		if err != nil {
			fmt.Println(err)
			return
		}

		result, err := calculate(operator, firstNumber, secondNumber)
		if err != nil {
			fmt.Println("this is a error: ", err)
			return
		}

		fmt.Printf("Result: %v %s %v = %v\n", firstNumber, operator, secondNumber, result)
	}
}

func calculate(operator string, num1 float64, num2 float64) (float64, error) {

	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("division by zero")
		}
		return num1 / num2, nil

	case "*":
		return num1 * num2, nil
	default:
		return 0, nil
	}
}

func getFirstNumber() float64 {
	var strNum1 string

	for {
		fmt.Println("Enter first number: ")
		_, _ = fmt.Scanln(&strNum1)
		number1, err := strconv.ParseFloat(strNum1, 64)
		if err != nil {
			fmt.Println("invalid Input: ", strNum1)
			//log.Println(err)dsd
			continue
		}

		return number1
	}

}

func getSecondNumber() (float64, error) {
	var strNum string

	for {
		fmt.Println("Enter second number: ")
		_, _ = fmt.Scanln(&strNum)
		number, err := strconv.ParseFloat(strNum, 64)
		if err != nil {
			fmt.Println("invalid Input: ", strNum)
			//log.Println(err)dsd
			continue
		}
		return number, nil
	}

}

func getOperator() string {
	var operator string

	for {
		fmt.Print("Enter operator (+, -, *, /): ")
		_, _ = fmt.Scanln(&operator)

		if operator == "+" || operator == "-" || operator == "*" || operator == "/" {
			return operator
		} else {
			fmt.Println("invalid operator")
			continue
		}
	}
}
