package main

import (
	"github.com/MicahParks/cereal"
)

var (

	// numbers is a alice of runes that holds all numbers.
	numbers = []rune("0123456789")
)

func main() {

	// Create a cereal generator.
	gen, _ := cereal.Generator(numbers, "")

	// Print the next cereal in the generator forever.
	for i := 0; i < 100; i++ {
		println(gen())
	}
}
