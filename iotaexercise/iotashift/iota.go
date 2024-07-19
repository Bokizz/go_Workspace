package main

import "fmt"

// iota e alat koj ni pomaga so definiranje ili grupiranje taka shto mozhe konstanti da definirame so iota posle nivnoto ime
const (
	c0 = iota // c0 == 0 prvata deklaracija na iota e sekogash 0 osven ako ima nekoja formula za presmetka
	c1 = iota // c1 == 1 dokolku imashe formula istata bi prodolzhila da se vrshi i za drugite promenlivi posle prvata iota inicijalizirana
	c2 = iota // c2 == 2
)

const (
	c3 = iota + 3 // c0 == 3 tuka e so formula
	c4            // c4 = 4 ,mozhe i bez da pisheme = iota zatoa shto samo ednash treba da se deklarira za edna promenliva i potoa site pod taa promenliva se isto iota
	c5            // c5 = 5
	c6            // c6 = 6
)

func main() {
	// fmt.Println(c0, c1, c2) Vaka prvo gi pechati i gi stava vo raspored po redosled
	// fmt.Println(c3, c4, c5, c6)
	// fmt.Printf("%d \t %b\n", 1, 1) 1 = c1
	// fmt.Printf("%d \t %b\n", 1<<1, 1<<1) za 1 mesto levo shiftiraj
	// fmt.Printf("%d \t %b\n", 1<<2, 1<<2) za 2
	// fmt.Printf("%d \t %b\n", 1<<3, 1<<3) za 3
	// fmt.Printf("%d \t %b\n", 1<<4, 1<<4) za 4
	// fmt.Printf("%d \t %b\n", 1<<5, 1<<5) za 5
	// fmt.Printf("%d \t %b\n", 1<<6, 1<<6) za 6
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<c1, 1<<c1)
	fmt.Printf("%d \t %b\n", 1<<c2, 1<<c2)
	fmt.Printf("%d \t %b\n", 1<<c3, 1<<c3)
	fmt.Printf("%d \t %b\n", 1<<c4, 1<<c4)
	fmt.Printf("%d \t %b\n", 1<<c5, 1<<c5)
	fmt.Printf("%d \t %b\n", 1<<c6, 1<<c6)
	// istoto se dobiva i so komentiraniot pristap
}
