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
func BenchmarkGenerator(b *testing.B) {

	// Create the cereal generator.
	gen, err := cereal.Generator(alphabet, "")
	if err != nil {
		b.Errorf("Failed to creat cereal generator.\nError: %s\n", err.Error())
		b.FailNow()
	}

	// The number of cereal generations to take place.
	iterations := 1000000

	// Generate the desired number of cereal.
	for i := 0; i < iterations; i++ {
		_ = gen()
	}
}
