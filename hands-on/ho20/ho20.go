package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 251
	y := rand.Intn(x)
	fmt.Println("Random number Y is ", y)
	if y >= 0 && y <= 100 {
		fmt.Println("Random number Y is between 0 and 100")
	} else if y >= 101 && y <= 200 {
		fmt.Println("Random number Y is between 101 and 200")
	} else if y >= 201 && y <= 250 {
		fmt.Println("Random number Y is between 201 and 250")
	} else {
		fmt.Println("Not in the scope")
	}
	for i := 0; i < 5; i++ {
		z := rand.Intn(3)
		fmt.Println(z)
	}
	k := rand.Intn(251)
	switch {
	case k <= 100:
		fmt.Println("Random number K is between 0 and 100")
	case k >= 101 && k <= 200:
		fmt.Println("Random number K is between 101 and 200")
	case k >= 201 && k <= 250:
		fmt.Println("Random number K is between 201 and 250")
	default:
		fmt.Println("Out of scope")
	}
}
