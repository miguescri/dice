package main

import (
	"fmt"
	"github.com/miguescri/dice"
)

func main() {
	sides := 6
	d, err := dice.NewDice(sides)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Roll:", d.Roll())
	fmt.Println("RollN(10):", d.RollN(10))

	r, rs := d.SumN(10)
	fmt.Println("SumN(10): ", r, rs)

	r, rs = d.SumNK(10, 5)
	fmt.Println("SumNK(10,5): ", r, rs)
}
