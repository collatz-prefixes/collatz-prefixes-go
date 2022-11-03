package utility

import (
	c "collatz-prefixes/internal/common"
	"math/big"
)

// Finds the path from a number.
func NTOP(n *big.Int) []bool {
	n.Sub(n, c.ONE) // decrement
	b := NTOB(n)    // to binary
	b = REVERSE(b)  // reverse
	b = FLIP(b)     // flip
	return b
}

func PTON(p []bool) *big.Int {
	// todo
}

// Given a list of bools as binary representation, returns the corresponding number.
func BTON(b []bool) *big.Int {

}

func NTOB(n *big.Int) []bool {
	return []bool{} // todo
}

// Flips the values in a boolean list.
func FLIP(b []bool) []bool {
	for i, b_i := range b {
		b[i] = !b_i
	}
	return b
}

func REVERSE(b []bool) []bool {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
