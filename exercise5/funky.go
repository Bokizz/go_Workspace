package main

import (
	"fmt"

	"example.com/go_Workspace/exercise5/funky2" // prvoto example.com/go_Workspace e imeto na mojot modul za zhal ostanato e patekata do funky2 package-ot
)

func main() {
	fmt.Println("Now we are going to take another external package and use it in this so called main package")
	funky2.Funky2() // ova e del od funky2 package koi shto e sosema poseben package
}
