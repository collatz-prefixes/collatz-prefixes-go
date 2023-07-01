package collatz_prefixes

import "math/big"

// Finds the next number that resides at the path of `n`.
//
// The path is also given, as `n` can be in different paths (see path extension).
func RiptreeNextInPath(n *big.Int, p []bool) *big.Int {
	return n.Add(n, new(big.Int).Exp(TWO, big.NewInt(int64(len(p))), nil))
}

// Finds the prefix of a number, or a number at the given path.
func RiptreePrefixFind(n *big.Int, p []bool) []uint {
	// TODO: does this always work for even larger paths?
	if PTON(p).Cmp(n) != 0 {
		panic("Path does not match the number")
	}

	if ISPOW2(n) {
		ans := uint(0)
		for ; n.Cmp(ONE) == 1; ans++ {
			n.Rsh(n, 1)
		}
		return []uint{ans}
	} else {
		return PrefixFind(n, RiptreeNextInPath(n, p))
	}
}
