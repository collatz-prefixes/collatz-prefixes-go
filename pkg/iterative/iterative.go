package iterative

import (
	"collatzprefixes/internal/common"
	"collatzprefixes/pkg/prefix"
	"collatzprefixes/pkg/utils"
	"fmt"
	"math/big"
)

// Find the ECF by iteratively extending the path until prefix iteration results in 1.
func PathExtension(n *big.Int, prefixFinder func(n *big.Int, p []bool) []uint) []uint {
	p := utils.ToPath(n)
	pf := prefixFinder(n, p)
	for prefix.Iterate(n, pf).Cmp(common.ONE) != 0 {
		fmt.Printf("PF:  %v\n", pf)
		pf = prefixFinder(n, p)
		p = append(p, true)
	}
	return pf
}

// Find the ECF by iteratively consuming the prefix until the iteration result is 1.
func Prefix(n *big.Int, prefixFinder func(n *big.Int, p []bool) []uint) []uint {
	n = common.Copy(n)
	ans := make([]uint, 0)
	for {
		pf := prefixFinder(n, utils.ToPath(n))
		fmt.Printf("ANS: %v\nCUR: %v\nN:   %v\n\n", ans, pf, n)
		ans = prefix.Add(ans, pf)
		n = prefix.Iterate(n, pf)
		if n.Cmp(common.ONE) == 0 {
			return ans
		} else {
			n.Mul(n, common.THREE).Add(n, common.ONE)
			if len(ans) != 0 {
				ans = append(ans, ans[len(ans)-1])
			}
		}
	}
}
