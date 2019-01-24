package main

import "fmt"


type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	JPY
)




func main() {
	symbol := [...]string{USD:"$", EUR: "&", GBP: "^", JPY:"@"}
	fmt.Println(JPY, symbol[JPY])
}
