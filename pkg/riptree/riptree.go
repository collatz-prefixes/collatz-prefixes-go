package riptree

import (
	"collatzprefixes/internal/common"
	"collatzprefixes/pkg/prefix"
	"collatzprefixes/pkg/utils"
	"math/big"
)

// Finds the next number that resides at the path of `n`.
//
// The path is also given, as `n` can be in different paths (see path extension).
func NextInPath(n *big.Int, p []bool) *big.Int {
	return new(big.Int).Add(n, new(big.Int).Lsh(common.ONE, uint(len(p))))
}

// Finds the prefix of a number, or a number at the given path.
//
// If you only care about the number, simply pass NTOP(n) as the path.
func PrefixFind(n *big.Int, p []bool) []uint {
	// TODO: does this always work for even larger paths?
	if utils.PTON(p).Cmp(n) != 0 {
		panic("Path does not match the number")
	}

	n = common.Copy(n)
	if utils.ISPOW2(n) {
		ans := uint(0)
		for ; n.Cmp(common.ONE) == 1; ans++ {
			n.Rsh(n, 1)
		}
		return []uint{ans}
	} else {
		return prefix.Find(n, NextInPath(n, p))
	}
}
