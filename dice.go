package dice

import (
	"errors"
	"math/rand"
	"sort"
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

// SumNK returns the sum of k highest results of n dice rolls, and the list of all the individual results.
func (d Dice) SumNK(n, k int) (int, []int) {
	rs := d.RollN(n)
	// If the number of dices is greater than the number to take, order the list to grab the highest
	if n > k {
		sort.Sort(sort.Reverse(sort.IntSlice(rs)))
	} else {
		k = n
	}

	s := 0
	for i := 0; i < k; i++ {
		s += rs[i]
	}
	return s, rs
}

// SumN returns the sum of n dice rolls, and the list of individual results.
func (d Dice) SumN(n int) (int, []int) {
	return d.SumNK(n, n)
}
