package main

import (
	"fmt"
)

func main() {
	// razlichni nachini kako mozhe da se koristi
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	y := 6
	for y < 10 {
		fmt.Println(y)
		y++
	}
	for {
		fmt.Println(y)
		if y == 16 {
			break
		}
		y++
	}

	for i := 0; i < 20; i++ {
		if i%2 != 0 { // interesting gi ripa tija shto imaa ostatok
			continue
		}
		fmt.Println("I is ", i)
	}
	for i := 0; i < 10; i++ { // pola od triagolnik ili ostar triagolniik
		var k int = i
		for j := 0; j < 10-i-1; j++ {
			switch k {
			case 1:
				fmt.Printf(" ")
				k = 0 // trgni k = 0 i kje izgleda pointeresna figura
			case 2:
				fmt.Printf("  ")
				k = 0
			case 3:
				fmt.Printf("   ")
				k = 0
			case 4:
				fmt.Printf("    ")
				k = 0
			case 5:
				fmt.Printf("     ")
				k = 0
			case 6:
				fmt.Printf("      ")
				k = 0
			case 7:
				fmt.Printf("       ")
				k = 0
			case 8:
				fmt.Printf("        ")
				k = 0
			case 9:
				fmt.Printf("         ")
				k = 0
			}
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	m := map[string]int{ // map go koristime za data strukturiranje taka shto davame kljuch promenliva so vrednost
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Println("Ranging over the map", k, v) // random map dava redosled
	}
}
