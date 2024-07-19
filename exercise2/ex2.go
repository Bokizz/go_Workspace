package main

import "fmt"

func main() {
	// short declaration
	a := 42
	fmt.Println(a)

	b, c, d := 43, 1, "Doom"
	fmt.Println(b+c, d)

	// declaration

	var g int
	fmt.Println(g) // dava 0 vrednost samo ako e deklarirano bez vrednost
	g = 43
	fmt.Println(g)
	var h int = 500 // deklarirano so inicijalizacija na vrednost
	fmt.Println(h)

	// ne raboti koga edna varijabla koja shto e deklarirana ne e iskoristena vo programata
}
