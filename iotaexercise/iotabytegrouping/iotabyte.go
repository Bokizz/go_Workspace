package main

import "fmt"

type ByteSize int // definirame nov tip
const (
	_           = iota
	KB ByteSize = 1 << (10 * iota) // sekoja naredna iteracija e 2^10 * 2^10 * 2^10.... vake samo pravime da se mnozhi stepenot celo vreme so iota shto ni krati muka i pravime shiftiranje soodvetno
	MB                             // zirni kako izgledaat MB GB TB KB vo bitski prikaz i kje ti bide jasno zashto se shiftira(za sekoe se stava 1 kec na x*10ti bit kade x raste za sekoja nova skala dostignata)
	GB
	TB
)

func main() {
	fmt.Printf(" KB = %d bits, \t \t binary = %b\n", KB, KB)
	fmt.Printf(" MB = %d bits, \t \t binary = %b\n", MB, MB)
	fmt.Printf(" GB = %d bits, \t \t binary = %b\n", GB, GB)
	fmt.Printf(" TB = %d bits, \t binary = %b\n", TB, TB)

}
