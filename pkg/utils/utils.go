package utils

import (
	common "collatzprefixes/internal/common"
	"math/big"
)

// Finds the path from a number.
func NTOP(n *big.Int) []bool {
	n = common.Copy(n)

	n.Sub(n, common.ONE) // decrement
	b := NTOB(n)         // to binary
	REVERSE(b)           // reverse
	FLIP(b)              // flip

	return b
}

// Finds the number from a path.
func PTON(p []bool) *big.Int {
	pp := make([]bool, len(p))
	copy(pp, p)

	FLIP(pp)             // flip
	REVERSE(pp)          // reverse
	n := BTON(pp)        // to decimal
	n.Add(n, common.ONE) // increment

	return n
}

// Given a list of bools as binary representation, returns the corresponding number.
func BTON(b []bool) *big.Int {
	n := big.NewInt(0)
	for _, b_i := range b {
		n.Lsh(n, 1)
		if b_i {
			n.Add(n, common.ONE)
		}
	}
	return n
}

// Convert a number to binary format.
func NTOB(n *big.Int) []bool {
	n = common.Copy(n)
	b := make([]bool, 0)
	for n.Cmp(common.ZERO) != 0 {
		b = append(b, n.Bit(0) == 1)
		n.Rsh(n, 1)
	}
	return REVERSE(b)
}

// Flips the values in a boolean list.
func FLIP(b []bool) []bool {
	for i, b_i := range b {
		b[i] = !b_i
	}
	return b
}

// Reverse the array of bits.
func REVERSE(b []bool) []bool {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

// Returns true if the number is a power of two.
func ISPOW2(n *big.Int) bool {
	if n.Cmp(common.ZERO) == 0 {
		return true
	} else {
		// n & (n-1) == 0
		return common.ZERO.Cmp(new(big.Int).And(n, new(big.Int).Sub(n, common.ONE))) == 0
	}
}

// Returns the next power of two for the given number.
func NEXTPOW2(n *big.Int) *big.Int {
	if ISPOW2(n) {
		return n
	}

	p := big.NewInt(1)
	for p.Cmp(n) == -1 {
		p.Lsh(p, 1)
	}

	return p
}
