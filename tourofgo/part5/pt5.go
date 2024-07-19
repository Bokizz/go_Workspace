package main

import "fmt"

const (
	Big   = 1 << 100  // shiftiranje na levo za 100 mesta
	Small = Big >> 99 // shiftiranje na desno za 99 mesta
)

func needInt(x int) int           { return (x*10 + 1) } // vrakja cel broj
func needFloat(x float64) float64 { return (x * 0.1) }  // vrakja float
// func needUnInt(x uint) uint       { return (x*10 + 1) }
func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needUnInt(Big)) ne mozhe ova bidejkji e pregolem broj
}
