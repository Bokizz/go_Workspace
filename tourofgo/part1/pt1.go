package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello the time is", time.Now(), " and your favourite random number is ", rand.Intn(1000))
	fmt.Println("Root of 9 is ", math.Sqrt(9))
}
