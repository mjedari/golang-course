package calculator

var Temp string
var privateTemp string

func Sum(x int, y int) int {
	return x + y
}

// Minus subtract two numbers
func Minus(x int, y int) int {
	return x - y
}

func divide(x, y int) int64 {
	return int64(x / y)
}

func Multiple(x, y int) int {
	return x * y
}
