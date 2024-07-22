package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

func init() { // init se koristi za odrzhuvanje na stabilen protok na izvrshuvanje kod
	fmt.Println("Begin initialization")
}

func main() {
	fmt.Println("First segment.")
	x := 6
	y := 7

	if x > y {
		fmt.Println("x is bigger than y")
	} else if x == y {
		fmt.Println("x is equal to y")
	} else {
		fmt.Println("x is less than y")
	}
	if x == (y-1) && y > 5 {
		fmt.Println("Eureka!")
	} else if x == (y-1) && y < 5 {
		fmt.Println("Not good...")
	}

	x = 3
	switch {
	case x < 2:
		fmt.Println("X is less than 2")
	case x > 2:
		fmt.Println("X is bigger than 2")
	default:
		fmt.Println("X is 2")
	}
	y = 5
	switch y {
	case 4:
		fmt.Println("Y  is 4")
	case 5:
		fmt.Println("Y  is 5")
	case 6:
		fmt.Println("Y  is 6")
	case 7:
		fmt.Println("Y  is 7")
	default:
		fmt.Println("No idea")
	}
}
