package cereal

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
func Generator(alphabet []rune) func() string {

	// Create the root cereal.
	latestCereal := &cereal{
		alphabet:       alphabet,
		alphabetLength: uint64(len(alphabet)),
	}

	// Create a closure that will serve as a cereal generator.
	return func() (str string) {
		latestCereal, str = latestCereal.next()
		return str
	}
}
