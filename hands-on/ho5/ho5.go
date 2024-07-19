package main

import "fmt"

func main() {
	var a int
	fmt.Printf("%d - type : %T\n", a, a)
	b := "Kaj si be"
	fmt.Printf("%s - type : %T\n", b, b)
	c, d, e, f := 2, 3, 4, 6
	var k float64 = 43.69696969
	fmt.Printf("%v - type : %T\n", c, c)
	fmt.Printf("%v - type : %T\n", d, d)
	fmt.Printf("%v - type : %T\n", e, e)
	fmt.Printf("%v - type : %T\n", f, f)
	fmt.Printf("%v - type : %T", k, k)

}
