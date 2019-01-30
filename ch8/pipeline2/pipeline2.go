package main

import "fmt"

func main()  {

	naturals := make(chan int, 5)
	squares := make(chan int, 5)

	// counter
	go func() {
		for x:=0; x <= 10; x++ {
			fmt.Println("Counter: " ,x)
			naturals <- x
		}
		close(naturals)
	}()


	//squarer
	go func() {
		for x := range naturals {
			fmt.Println("Squarer:", x*x)
			squares <- x*x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x:= range squares {
		fmt.Println("main:", x)
	}
}
