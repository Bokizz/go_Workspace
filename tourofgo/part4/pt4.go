package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
		 tipovi: complex 64  complex128 ,
	 uint uint8 uint16 uint32 uint64 uintptr (unsigned integer),
	 int int8 int16 int32 int64
	 rune e tip na promenliva int32
	 byte e tip na promenliva uint8
	 unsigned integer gi zema samo pozitivnite broevi od 0 pa nagore int zema i negativni

broevite koi doagjaat posle imeto na tipot se opsegot do 2^x kade shto x e toj broj
*/
var (
	Tobe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T, Value: %v\n", Tobe, Tobe)
	fmt.Printf("Type: %T, Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T, Value: %v\n", z, z)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// var k uint = f vaka ne
	var k uint = uint(f) // vaka da
	fmt.Println(x, y, k)
}
