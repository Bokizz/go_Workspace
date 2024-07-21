package main

import "fmt"

var x float64 = 69.0

const (
	y = iota + 3/4
	b
	c
)

func main() {
	s := "This is just a exercise"
	fmt.Println(y, " ", b, " ", c)
	fmt.Println(s)
	fmt.Println("What is his power?")
	a := x * 130.434782609
	fmt.Printf("Its over %F!!!", a)
}
