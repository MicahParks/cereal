package cereal_test

import (
	"testing"

	"github.com/MicahParks/cereal"
)

var (

	// alphabet is a slice of runes that holds all lowercase letters.
	alphabet = []rune("abcdefghijklmnopqrstuvwxyz")
)

// BenchmarkGenerator benchmarks the cereal generator implementation.
func BenchmarkGenerator(_ *testing.B) {

	// Create the cereal generator.
	gen := cereal.Generator(alphabet)

	// The number of cereal generations to take place.
	iterations := 1000000

	// Generate the desired number of cereal.
	for i := 0; i < iterations; i++ {
		_ = gen()
	}
}
