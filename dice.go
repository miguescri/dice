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

// New creates a new Dice given a positive number of sides.
// If sides is non-positive, returns an error.
func New(sides int) (Dice, error) {
	if sides <= 0 {
		return Dice{}, errors.New("tried to make a dice with a non-positive number of sides")
	}
	return Dice{sides}, nil
}

// Roll returns a random number in the interval [1, d.sides]
func (d Dice) Roll() int {
	return rand.Intn(d.sides) + 1
}

// RollN returns the results of n dice rolls. If n is non-positive, returns an empty list
func (d Dice) RollN(n int) []int {
	if n <= 0 {
		return []int{}
	}
	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = d.Roll()
	}
	return l
}

// SumNK returns the sum of k highest results of n dice rolls, and the list of all the individual results.
// If k is greater than n, assumes k is n.
// If n or k are non-positive, assumes they are zero.
func (d Dice) SumNK(n, k int) (int, []int) {
	if n <= 0 {
		return 0, []int{}
	}

	rs := d.RollN(n)

	if k <= 0 {
		return 0, rs
	}

	var rsc []int // Pointer to the ordered list to sum
	// Order the list if there are more dices than results to sum
	if n > k {
		// Create a copy of the results
		rsc = make([]int, n)
		copy(rsc, rs)
		// Order the list from high to low
		sort.Sort(sort.Reverse(sort.IntSlice(rsc)))
	} else {
		rsc = rs
		// Set n as the loop limit
		k = n
	}

	s := 0
	for i := 0; i < k; i++ {
		s += rsc[i]
	}
	return s, rs
}

// SumN returns the sum of n dice rolls, and the list of individual results.
func (d Dice) SumN(n int) (int, []int) {
	return d.SumNK(n, n)
}
