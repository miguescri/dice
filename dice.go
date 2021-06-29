package dice

import (
	"errors"
	"math/rand"
)

type Dice struct {
	sides int
}

// NewDice creates a new Dice given a positive number of sides
func NewDice(sides int) (Dice, error) {
	if sides <= 0 {
		return Dice{}, errors.New("tried to make a dice with a non-positive number of sides")
	}
	return Dice{sides}, nil
}

func (d Dice) Roll() int {
	return rand.Intn(d.sides) + 1
}

func (d Dice) RollN(n int) []int {
	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = d.Roll()
	}
	return l
}

func (d Dice) SumN(n int) (int, []int) {
	rs := d.RollN(n)
	s := 0
	for _, r := range rs {
		s += r
	}
	return s, rs
}
