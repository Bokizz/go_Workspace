package main

import "fmt"

func main() {
	z := 42.5 // float64
	fmt.Println(z)
	y := 42.0      // float64
	fmt.Println(y) // dava 42ne e float togash
	var m float32 = 43.768889
	fmt.Printf("m - %v of type %T\n", m, m)
	// y = m // ne mozhe ova bidejki y e float64 a m e float32
	// fmt.Printf("%v of type %T", y, y)

	var x float64 = 50.542424242
	y = x
	fmt.Printf("y,x - %v of type %T", y, y) // od ist tip se pa raboti
}
