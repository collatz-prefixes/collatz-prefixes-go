package collatz

import "math/big"

// Finds the path from a number.
func NTOP(n *big.Int) []bool {
	n.Sub(n, ONE) // decrement
	b := NTOB(n)  // to binary
	REVERSE(b)    // reverse
	FLIP(b)       // flip
	return b
}

// Finds the number from a path.
func PTON(p []bool) *big.Int {
	FLIP(p)       // flip
	REVERSE(p)    // reverse
	n := BTON(p)  // to decimal
	n.Add(n, ONE) // increment
	return n
}

// Given a list of bools as binary representation, returns the corresponding number.
func BTON(b []bool) *big.Int {
	n := big.NewInt(0)
	for _, b_i := range b {
		n.Lsh(n, 1)
		if b_i {
			n.Add(n, ONE)
		}
	}
	return n
}

// Convert a number to binary format.
func NTOB(n *big.Int) []bool {
	b := make([]bool, 0)
	for n.Cmp(ZERO) != 0 {
		// prepend n & 1 to the beginning
		b = append(b, n.Bit(0) == 1)
		n.Rsh(n, 1)
	}
	REVERSE(b) // we want the LSB to be the first index
	return b
}

// Flips the values in a boolean list.
func FLIP(b []bool) {
	for i := range b {
		b[i] = !b[i]
	}
}

// Reverse the array of bits.
func REVERSE(b []bool) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

// Returns true if the number is a power of two.
func ISPOW2(n *big.Int) bool {
	if n.Cmp(ZERO) == 0 {
		return true
	} else {
		n_1 := new(big.Int).Sub(n, ONE)
		return ZERO.Cmp(new(big.Int).And(n, n_1)) == 0
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