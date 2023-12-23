package main

import (
	"fmt"
	"log"
)

func main() {
	var number, number2 float64
	n, err := fmt.Scanf("%f %f", &number, &number2)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("this is number of reads: ", n, number, number2)
}

/*
* 	1. if we put %d in fmt.Scanf() we can expect the error if the user do not provide int input
	2. n is not the number of bytes in fmt.Scanf(). its simple the number of reads that happened in the provided format.
	3.

*/
