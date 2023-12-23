package main

import "fmt"

func main() {

	myMap := make(map[string]string)
	//var myMap map[string]string

	myMap["hesam"] = "27"
	myMap["mehdi"] = "33"

	fmt.Println("this is my Map:", myMap)

	// update
	myMap["mehdi"] = "35"

	fmt.Println("this is my Map:", myMap)

	// delete
	delete(myMap, "mehdi")
	fmt.Println("this is my Map:", myMap)

	member, ok := myMap["mehdi"]
	if ok {
		fmt.Println("this is member ", member)
	}
}
