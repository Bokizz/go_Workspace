package main

import "fmt"

// koga ne se deklarirani so specifichna vrednost ja dobivaat 0 vrednosta soodvetno( za bool false )
var c, python, java bool
var j, z int = 1, 2

const d int = 42

func main() {
	var i int
	c, python, java = true, false, true // override
	fmt.Println(i, c, python, java, d)
	fmt.Println(j, z)
	k := 3       // koga vake definirame kompajlerot sam doznava shto tip e promenlivata
	var p = 4.23 // isto i tuka sam doznava kombajlerot
	fmt.Printf("k is %T p is %T", k, p)
}
