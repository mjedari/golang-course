package main

import (
	"fmt"
	"strconv"
	"time"
)

//

//func myFunc(str string) {
//	// anything
//	fmt.Println("I'm myFunc")
//}
//
//func myMainFunc(f func(str string)) {
//	f("some str")
//	fmt.Println("I'm myMainFunc")
//}

// return number
func power(i interface{}) int {
	str, ok := i.(string)
	if !ok {
		return 0
	}

	m, _ := strconv.Atoi(str)
	return m * m
}

type MyInt int

func main() {

	//myMainFunc(myFunc)

	// assertion
	//num := 10
	//value := power(num)
	//
	//fmt.Printf("this is power of %d: %d \n", num, value)

	// conversion

	//i := 10
	//mi := MyInt(10)
	//fmt.Printf("type of i %T \n", i)
	//fmt.Printf("type of mi %T \n", int(mi))

	//f := func() {
	//	for {
	//		time.Sleep(time.Second)
	//		fmt.Println("This is a function")
	//	}
	//}
	//
	//go f()
	//
	//for {
	//	time.Sleep(time.Second)
	//	fmt.Println("This is the end of the program")
	//}

	// Channel
	// unbuffered channel when send happens, there should be some area to read that channel unless we should expect deadlock
	// buffered channel can accept send exact amount of buffer size without necessarily waiting to be received data

	//var ch2 chan string
	ch := make(chan string)
	// [ [ _string,  _string ]
	//ch3 := make(chan []string, 2)
	// [ [ _[]string,  _[]string ]

	go listenChannel(ch)

	ch <- "mahdi" // <- deadlock here

	//close(ch)

	//for {
	//	time.Sleep(time.Second)
	//	fmt.Println(<-ch)
	//}

	// select version
	//select {
	//case v := <-ch:
	//	fmt.Println(v)
	//default:
	//	return
	//}
	time.Sleep(3 * time.Second)
}

func listenChannel(ch chan string) {
	fmt.Println("read data from channel: ", <-ch)
}
