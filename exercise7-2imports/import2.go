package main

import (
	"fmt"

	"github.com/Bokizz/temporary"  // sekogash iako e dodadeno vo module-ot
	"github.com/Bokizz/temporary2" // za da se pristapi do funkcii koi ne se koristat kaj temporary iako e importiran temporary2 a gi ima kj temporary2
)

func main() {
	fmt.Println("The text below is from an imported repo")
	s1 := temporary.Hey() // mozhe da inicijalizirame
	fmt.Println(s1)
	fmt.Println(temporary.Listen()) // i direktno da go printame
	s2 := temporary.Listenedsom()
	fmt.Println(s2) // string e pa mora da se inicijalizira pa da se printa funkcijata
	s3 := temporary.Listenednon()
	fmt.Println("Second time eavesdropping.", s3)
	temporary.FromV11()
	temporary.FromV12()
	temporary2.FromV100()
	fmt.Println("Great success!")
}
