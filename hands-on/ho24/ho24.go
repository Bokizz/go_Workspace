package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 10
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("Turn ", i)
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Println("X is ", x, "; Y is ", y)
	}

	// for with condition (kao while)
	for x > 0 {
		fmt.Println(x)
		x--
	}

	// infinite loop so prekin break uslov
	for {
		fmt.Println(x)
		x++
		if x == 20 {
			break
		}
	}

	// odd numbers to 100 for loop
	for i := 0; i < 100; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// nested loop example
	for i := 0; i < 9; i++ {
		fmt.Printf("Layer %d:", i)
		var k int = i
		for j := i + 2; j < 20-i; j++ { // the upside down triangle :)
			for k != 0 {
				fmt.Printf(" ")
				k--
			}
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}
