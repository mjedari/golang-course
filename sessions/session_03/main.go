package main

import (
	"errors"
	"fmt"
)

func main() {

	// bool, in, float, string
	// slice (), array (array) , map (dict)

	//var a1 [4]string
	//s1 := make([]string, 5)

	// 0, 1, 2 , 3
	//[1, 2, 4, 5] => [4]int // array with 4 elements
	//myFunction("test", 5)

	// array
	//var a1 [4]int // 1.allocate
	//a1[0] = 200   // 2.assignment
	//a1[1] = 100
	//a1[2] = 200
	//a1[3] = 300
	//fmt.Println("value of a1: ", a1)
	//
	//// := initiate
	//a2 := [4]int{1, 2, 3, 4} // => assignment {}
	//fmt.Println("value of a2: ", a2)

	// slice
	//var s1 []int // 1.allocate => nil
	//fmt.Println("s1 is: ", s1)
	//s1[0] = 200  // 2.assignment
	//s1[1] = 100
	//s1[2] = 200
	//s1[3] = 300
	//fmt.Println("value of s1: ", s1)

	//// append
	//var s2 []int
	//s3 := append(s2, 200)
	//s4 := append(s3, 400)
	//fmt.Printf("value of s2: %v s3: %v s4: %v \n", s2, s3, s4)

	// make
	//var myList = make([]int, 4)
	//myList := make([]int, 4) // => [0, 0, 0, 0]
	//fmt.Println("my list is: ", myList)
	//myList[0] = 400
	//fmt.Println("value of myList: ", myList)

	// initiate
	//myNewList := []int{1, 2, 3, 4, 5, 6}
	//fmt.Println("value of myList: ", myNewList)
	//// add 7 to slice
	//myNewList = append(myNewList, 7)
	//fmt.Println("value of myNewList after appending: ", myNewList)

	// loop
	// for => 3 model
	// for range

	// for i
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// for (condition)
	//var state bool
	//state = true

	//for state {
	//	fmt.Println("HI")
	//
	//	// continue
	//	continue
	//	fmt.Println("this is not goiing to print")
	//
	//}

	//for state {
	//	fmt.Println("HI")
	//	break
	//}

	//students := make([]string, 0)
	students := []string{"hadi", "reza"}

	for _, value := range students {
		fmt.Println(value)
	}
	fmt.Println("lenght: ", len(students))

	// how to check student is empty or not
	if len(students) > 1 {
		fmt.Println("students is not empty")
	} else {
		fmt.Println("students is empty")
	}

	err := ageChecker(17)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("end of the line")

}

func ageChecker(age uint) error {
	if age < 18 {
		return errors.New("the age is under 18")
	}

	return nil
}

func myFunction(myArguman string, age uint) {

}
