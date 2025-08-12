package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.SetString("9999999999999999", 10)

	b := new(big.Int)
	b.SetString("999999999999999", 10)

	a = a.Mul(a, b)
	fmt.Printf("%s\n", a.String())

	a = a.Div(a, b)
	fmt.Printf("%s\n", a.String())

	a = a.Sub(a, b)
	fmt.Printf("%s\n", a.String())

	a = a.Add(a, b)
	fmt.Printf("%s\n", a.String())
}
