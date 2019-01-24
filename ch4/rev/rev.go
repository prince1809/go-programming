package main

import "fmt"

func main() {

	a := [...]int{0, 1, 2, 3, 4, 5}
	//x := a
	//printReverse(x[:])
	z := a
	rotateByNElements(2, z[:])
	//fmt.Println(a)
}

func rotateByNElements(n int, s []int) {
	fmt.Println(s)  // [0 1 2 3 4 5]
	reverse(s[:n])  // [1 0 2 3 4 5]
	fmt.Println(s)
	reverse(s[n:]) // [1 0 5 4 3 2]
	fmt.Println(s)
	reverse(s)  // [2 3 4 5 0 1]
	fmt.Println(s)
}

func printReverse(s [] int) {
	reverse(s[:])
	fmt.Println(s)
}

// reverse reverses a slice of ints in place
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
