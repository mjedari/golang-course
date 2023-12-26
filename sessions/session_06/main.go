package main

import (
	"fmt"
	"sync"
)

// lesson 1
//var count int
//
//func main() {
//
//	for i := 0; i < 10; i++ {
//		//count++
//		go process(i) // spawn
//	}
//
//	select {}
//}
//
//func process(count int) {
//	fmt.Println(count)
//}

// lesson 2

//func main() {
//	number := 1
//	for i := 0; i < 10; i++ {
//		go process(&number) // go routine spawns here
//	}
//
//	fmt.Println("This is the number: ", number)
//	time.Sleep(3 * time.Second)
//}
//
//func process(number *int) {
//	*number++
//}

/*
	lesson 3
*/
// 1 tread
func main() {

	var wg sync.WaitGroup
	// [1,1, 1, 1, 1, 1, 1, 1, 1]
	for i := 0; i < 10; i++ {
		wg.Add(1)

		//go process(i) // spawn

		// approach 1
		//go func() {
		//	processV1(i)
		//	wg.Done()
		//}()

		// approach 2
		go processV2(i, &wg)
	}

	wg.Wait() // 10 members
	// continue => 0 member
	fmt.Println("End of program")
}

func processV1(i int) {
	fmt.Printf("this is the %dth process\n", i)
}

func processV2(i int, wg *sync.WaitGroup) {
	fmt.Printf("this is the %dth process\n", i)
	wg.Done()
}
