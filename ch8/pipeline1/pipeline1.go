package main

import "fmt"

func main() {
	nat := make(chan int)
	sq := make(chan int)

	// counter
	go counter(nat)

	//squares
	go squarer(nat, sq)

	for {
		fmt.Println(<-sq)
	}
}

func counter(naturals chan int) {
	for x := 0; ; x++ {
		naturals <- x
	}
}

func squarer(natural chan int, square chan int) {
	for {
		x := <-natural
		square <- x * x
	}
}
