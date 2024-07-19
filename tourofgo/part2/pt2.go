package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) { //deklariranjeto na promenlivi tuka e sledno: ime na promenliva pa tip na kraj od promenlivite ili posle sekoja promenliva posebno, posle zagradata shto ochekuvame da dobieme na kraj od funkcijata odnosno kakov tip na rezultat
	return y, x
}

func split(sum int) (x, y int) { // tuka pravime 'naked return' kade shto od start gi proglasuvame promenlivite koi gi vrakjame kako rezultat od funkcijata, tuka tie se x i y
	x = sum * 4 / 9
	y = sum - x
	return // sega ovoj return vaka samo pishan gi vrakja x i y
}

func main() {
	fmt.Println(add(42, 13))
	sayHello()
	a := add(100, 200)
	fmt.Println(a)

	z, b := swap("Hello", "Dude")
	fmt.Println(z, b)
	fmt.Println(`
	What
	The 
	Hell
	`)

	fmt.Println(split(17))
}

func sayHello() {
	fmt.Println("Hello there")
}
