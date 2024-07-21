package main

import (
	"fmt"

	"github.com/Bokizz/temporary"
)

func main() {
	fmt.Println("The text below is from an imported repo")
	s1 := temporary.Hey() // mozhe da inicijalizirame
	fmt.Println(s1)
	fmt.Println(temporary.Listen()) // i direktno da go printame
	s2 := temporary.Heardsom()
	fmt.Println(s2)
	fmt.Println(temporary.Heardnon())
}
