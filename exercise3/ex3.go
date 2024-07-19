package main

import "fmt"

//binary to hexadecimal
func main() { // so printf raboti dodavanje na vrednost na %b i %x inaku so println nema dodeluvanje na vrednosti tuku samo se pechatat na odredeno mesto
	a, b, c, d, e, f, g := 0, 1, 2, 3, 4, 5, 66
	fmt.Printf("%b to hexa => %x\n,%b to hexa => %x\n, %b to hexa => %x\n, %b to hexa => %x\n, %b to hexa => %x\n , %b to hexa => %x\n, real value %v => binary %b to hexa => %x\n", a, a, b, b, c, c, d, d, e, e, f, f, g, g, g)
}
