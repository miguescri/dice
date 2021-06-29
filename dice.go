package dice

import (
	"errors"
	"math/rand"
)

// Dice represents a dice with a given number of sides.
type Dice struct {
	sides int
}

// NewDice creates a new Dice given a positive number of sides.
// If sides is non-positive, returns an error.
func NewDice(sides int) (Dice, error) {
	if sides <= 0 {
		return Dice{}, errors.New("tried to make a dice with a non-positive number of sides")
	}
	return Dice{sides}, nil
}

// Roll returns a random number in the interval [1, d.sides]
func (d Dice) Roll() int {
	return rand.Intn(d.sides) + 1
}

// RollN returns the results of n dice rolls.
func (d Dice) RollN(n int) []int {
	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = d.Roll()
	}
	return l
}

// SumN returns the sum of n dice rolls, and the list of individual results.
func (d Dice) SumN(n int) (int, []int) {
	rs := d.RollN(n)
	s := 0
	for _, r := range rs {
		s += r
	}
	return s, rs
}
