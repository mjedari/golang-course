package main

import "fmt"

func main() {
	// concatinate
	//test := "name: " + "Hesam" + " " + "Keshavarz" + " "
	//fmt.Println(test)

	//output := fmt.Sprintf("Name: %s Family: %s", "Hesam", "Keshavarz")
	//fmt.Println(output)

	// get input from terminal

	// 1. define a variable
	//var name string

	// 2. get the variable
	//fmt.Println("Enter your name:")
	//fmt.Println("&name=", &name)
	//fmt.Scanln(&name)
	//
	//fmt.Println("Your name is: ", name)

	//var first, last string
	//fmt.Println("Enter your full name:")
	//fmt.Scanf("%s %s", &first, &last)
	//
	//fmt.Printf("Your first is: %s and last is: %s \n", first, last)

	for {
		var first, last string
		fmt.Println("Enter your full name:")
		fmt.Scanf("%s %s", &first, &last)

		fmt.Printf("Your first is: %s and last is: %s \n", first, last)
	}
}
