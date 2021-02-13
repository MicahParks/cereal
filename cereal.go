// Package cereal provides a generator of unique strings based on an alphabet in sequential order. When a cycle of
// the alphabet is complete, a new character is appended, and the cycle restarts.
//
// This is the first 21 strings returned from the generator when the alphabet `0123456789` is used.
//   0
//   1
//   2
//   3
//   4
//   5
//   6
//   7
//   8
//   9
//   00
//   01
//   02
//   03
//   04
//   05
//   06
//   07
//   08
//   09
//   10
//   11
package cereal

import (
	"errors"
)

var (

	// ErrBadStartingString indicates that the given starting string wasn't valid for the give alphabet.
	ErrBadStartingString = errors.New("couldn't create cereal generator for the given alphabet with the given starting string")

	// ErrBadAlphabet indicates that the given alphabet couldn't be used to make a cereal generator.
	ErrBadAlphabet = errors.New("couldn't create a cereal generator for the given alphabet (may contain duplicate rune)")
)

// cereal represents an index in the returned string of a cereal generation.
type cereal struct {
	alphabet       []rune
	alphabetIndex  uint64
	cerealIndex    uint64
	alphabetLength uint64
	parent         *cereal
}

// buildStr builds the returned string for a cereal generation.
func (c *cereal) buildStr() (str string) {

	// If there is parent, start the base of the returned string with their returned string.
	if c.parent != nil {
		str = c.parent.buildStr()
	}

	// Add the character at the current index.
	return str + string(c.alphabet[c.alphabetIndex])
}

// increment increments the current index of the cereal. It will increment its parent's index if it cycles. The root
// cereal will return true for completedCycle should it cycle. That value will be carried to the top of the recursive
// call.
func (c *cereal) increment() (completedCycle bool) {

	// Increment the index.
	c.alphabetIndex++

	// If the index is larger than the maximum index.
	if c.alphabetIndex == c.alphabetLength {

		// Reset the index to 0.
		c.alphabetIndex = 0

		// If there is a parent, increment it. If this is the root, propagate up the recursive stack a true value.
		if c.parent != nil {
			completedCycle = c.parent.increment()
		} else {
			completedCycle = true
		}
	}

	return completedCycle
}

// next generates the next return string for the cereal. It will return the newest cereal to call next() on later.
func (c *cereal) next() (latestCereal *cereal, str string) {

	// Assume the current cereal is the newest cereal.
	latestCereal = c

	// Create the return string.
	str = c.buildStr()

	// If the root cereal has completed a cycle, it's time to add a new cereal.
	if c.increment() {

		// Create a new cereal and mark it as the latest cereal.
		latestCereal = &cereal{
			alphabet:       c.alphabet,
			cerealIndex:    c.cerealIndex + 1,
			alphabetLength: c.alphabetLength,
			parent:         c,
		}
	}

	return latestCereal, str
}

// Generator creates a function that will return a cereal generator via a closure.
func Generator(alphabet []rune, starting string) (cerealGenerator func() string, err error) {

	// Create a block of code that will cause variables in this scope to get collected by GC when it's over.
	{
		// Create a map that will serve as a set.
		alphaSet := make(map[rune]struct{})

		// Iterate through the given alphabet. If there are any duplicates, return an error.
		var ok bool
		for _, r := range alphabet {
			if _, ok = alphaSet[r]; ok {
				return nil, ErrBadAlphabet
			}
			alphaSet[r] = struct{}{}
		}
	}

	// Determine the length of the alphabet.
	alphabetLength := uint64(len(alphabet))

	// Create the starting string.
	var latestCereal *cereal
	if starting == "" {

		// Create the root cereal.
		latestCereal = &cereal{
			alphabet:       alphabet,
			alphabetLength: alphabetLength,
		}
	} else {

		// Turn the starting string into a slice of runes.
		startingSlice := []rune(starting)

		// Iterate through the starting string's rune slice.
		var found bool
		for index, r := range startingSlice {

			// Keep track if this character is in the alphabet.
			found = false

			// Check if this character in the starting string is in the alphabet.
			for alphaIndex, alpha := range alphabet {
				if r == alpha {

					// If this character is in the starting string for the alphabet, create a cereal for it.
					cer := &cereal{
						alphabet:       alphabet,
						alphabetIndex:  uint64(alphaIndex),
						cerealIndex:    uint64(index),
						alphabetLength: alphabetLength,
						parent:         latestCereal,
					}

					// Make this current cereal the latest cereal.
					latestCereal = cer

					// Mark this character in the starting string as found.
					found = true
					break
				}
			}

			// If a character in the starting string is not in the alphabet, return an error.
			if !found {
				return nil, ErrBadStartingString
			}
		}
	}

	// Create a closure that will serve as a cereal generator.
	return func() (str string) {
		latestCereal, str = latestCereal.next()
		return str
	}, nil
}
