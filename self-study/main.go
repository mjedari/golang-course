package main

import "log"

func main() {
	defer log.Println("end of the function")

	log.Println("first lone of function")
	// line 1
	// line 2
	// line 3
	// ...
	// end
}
