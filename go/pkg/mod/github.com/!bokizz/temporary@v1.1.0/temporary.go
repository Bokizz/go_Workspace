package temporary

import (
	"fmt"

	"github.com/Bokizz/temporary2"
)

func Hey() string {
	return "Hey,hey!"
}

func Listen() string {
	return "Listen,hey!"
}
func Heardnon() string {
	return "Nothing of importance :)"
}
func Heardsom() string {
	return "They are taking the hobbits to Isengard"
}
func Listenednon() string {
	return temporary2.WhenListen(Heardnon())
}
func Listenedsom() string {
	return temporary2.WhenListen(Heardsom())
}

func FromV11() {
	fmt.Println("This is a text from version 1.1.0")
}
