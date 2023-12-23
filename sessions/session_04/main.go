package main

import (
	"fmt"
	"log"
)

func main() {
	// every thing is pass by value

	//input := "Hesam" // memory: 0x14000010240
	//fmt.Println("Address of input", &input)
	//
	//pointerInput := &input
	//fmt.Println("Address of input", *pointerInput)

	// pass pointer
	number := 10

	Increment(&number)

	fmt.Println(number)
}

func Increment(input *int) {
	*input++
}

func MyFunc() {
	defer log.Println("end of the function")
	log.Println("first line of function")
}
