package collatz

import (
	"math/big"
)

// Find the ECF by iteratively extending the path until prefix iteration results in 1.
func IterativePathExtension(n *big.Int, prefixFinder func(n *big.Int, p []bool) []uint) []uint {
	p := NTOP(n)
	for PrefixIterate(n, prefixFinder(n, p)).Cmp(ONE) != 0 {
		p = append(p, true)
	}
	return prefixFinder(n, p)
}

// Find the ECF by iteratively consuming the prefix until the iteration result is 1.
func IterativePrefix(n *big.Int, prefixFinder func(n *big.Int, p []bool) []uint) []uint {
	var pf []uint

	n = Copy(n)
	ans := make([]uint, 0)
	for {
		pf = prefixFinder(n, NTOP(n))
		ans = PrefixAdd(ans, pf)
		n = PrefixIterate(n, pf)
		if n.Cmp(ONE) == 0 {
			return ans
		} else {
			n.Mul(n, THREE).Add(n, ONE)
			if len(ans) != 0 {
				ans = append(ans, ans[len(ans)-1])
			}
		}
	}
}
