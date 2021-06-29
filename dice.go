package dice

import (
	"errors"
	"math"
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

// SumN returns the sum of n dice rolls, and the list of individual results.
func (d Dice) SumN(n int) (int, []int) {
	return d.SumNK(n, n)
}

// SumNK returns the sum of k highest results of n dice rolls, and the list of all the individual results.
// If k is greater than n, assumes k is n.
// If n or k are non-positive, assumes they are zero.
func (d Dice) SumNK(n, k int) (int, []int) {
	if n <= 0 {
		return 0, []int{}
	}
	rs := d.RollN(n)
	s := sum(rs, k)
	return s, rs
}

// Probability returns the probabilities of to obtain each value in [1, d.sides*k] when running SumNK(n,k)
func (d Dice) Probability(n, k int) []float64 {
	if n <= 0 || k <= 0 {
		return []float64{}
	}
	max := d.sides * k                              // Maximum value of SumNK(n,k)
	rolls := math.Pow(float64(d.sides), float64(n)) // Number of combinations of n dices
	counts := make([]float64, max)                  // Number of times each sum appears
	rs := make([]int, n)                            // Buffer to simulate the dice combinations

	prob(rs, d.sides, k, n-1, counts)

	// Calculate percentage for each combination
	for i := 0; i < max; i++ {
		counts[i] = counts[i] / rolls
	}
	return counts
}

// prob recursively generates every combination of n dices, calculates its sum and updates its count
func prob(rs []int, sides, k, n int, counts []float64) {
	if n >= 0 {
		for v := 1; v < sides+1; v++ {
			rs[n] = v
			prob(rs, sides, k, n-1, counts)
		}
	} else {
		s := sum(rs, k)
		counts[s-1] = counts[s-1] + 1
	}
}

func sum(rs []int, k int) int {
	if k <= 0 {
		return 0
	}
	n := len(rs)

	var rsc []int // Pointer to the copy of the list
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
	return s
}
