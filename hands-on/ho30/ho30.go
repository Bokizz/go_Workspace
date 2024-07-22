package main

import "fmt"

func main() {

	//simple data structure
	x := []int{42, 43, 44, 45, 46, 47}
	fmt.Println("Data structure")
	for i, v := range x {
		fmt.Printf("Index %v \t Value %v\n", i, v)
		i++
	}
	fmt.Println("")
	// data structure with key and value
	x1 := map[string]int{
		"James":       23, // key = James , value = 23
		"Pennysworth": 24,
	}
	fmt.Println("Data structure with key and value:")
	for k, v := range x1 {
		fmt.Printf("Key %v \t Value %v\n", k, v)
	}
}
