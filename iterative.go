package collatz_prefixes

import "math/big"

func IterativePathExtension(n *big.Int, prefixFinder func(p []bool) []uint) []uint {
	p := NTOP(n)
	pf := prefixFinder(p)
	for PrefixIterate(n, pf).Cmp(ONE) != 0 {
		p = append(p, true)
		pf = prefixFinder(p)
	}
	return pf
}

func IterativePrefix(n *big.Int, prefixFinder func(n *big.Int) []uint) []uint {
	ans := make([]uint, 0)
	pf := make([]uint, 0)
	for {
		pf = prefixFinder(n)
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
