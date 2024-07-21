package temporary2

import (
	"fmt"
	"strings"
)

func WhenListen(s string) string {
	return "When he listened he heard: " + strings.ToUpper(s)
}

func FromV100() {
	fmt.Println("This is text from temp2's version 1.0.0")
}
