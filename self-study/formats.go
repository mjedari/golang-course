package main

import (
	"fmt"
)

func main() {
	// Integer
	number := 42
	fmt.Printf("%%d: %d\n", number) // Base 10 integer
	fmt.Printf("%%b: %b\n", number) // Binary
	fmt.Printf("%%x: %x\n", number) // Hexadecimal (lowercase)
	fmt.Printf("%%X: %X\n", number) // Hexadecimal (uppercase)

	// Floating-point
	floatNum := 3.14159
	floatNum2 := 100000000000.2121
	fmt.Printf("%%f: %f\n", floatNum)  // Default float
	fmt.Printf("%%e: %e\n", floatNum2) // Scientific notation (lowercase)
	fmt.Printf("%%E: %E\n", floatNum2) // Scientific notation (uppercase)

	// String
	str := "hello"
	fmt.Printf("%%s: %s\n", str) // String
	fmt.Printf("%%q: %q\n", str) // Quoted string

	// Boolean
	boolVal := true
	fmt.Printf("%%t: %t\n", boolVal) // Boolean

	// Default format
	fmt.Printf("%%v: %v\n", number)   // Default format (integer)
	fmt.Printf("%%v: %v\n", floatNum) // Default format (float)
	fmt.Printf("%%v: %v\n", str)      // Default format (string)
	fmt.Printf("%%v: %v\n", boolVal)  // Default format (boolean)

	// Type
	fmt.Printf("%%T: %T\n", number)   // Type of variable (integer)
	fmt.Printf("%%T: %T\n", floatNum) // Type of variable (float)
	fmt.Printf("%%T: %T\n", str)      // Type of variable (string)
	fmt.Printf("%%T: %T\n", boolVal)  // Type of variable (boolean)

	// Pointer
	fmt.Printf("%%p: %p\n", &number) // Address of variable

	// Unicode character
	fmt.Printf("%c\n", 65)  // ASCII value for 'A'
	fmt.Printf("%c\n", 945) // Unicode value for Greek letter alpha (α)

	// Structs
	type Person struct {
		Name string
		Age  int
	}

	person := Person{"Hesam", 25}
	fmt.Printf("%+v\n", person) // useful for debugging, shows the struct key-values
	fmt.Printf("%#v\n", person) // useful for debugging, shows the struct type and key-values
}
