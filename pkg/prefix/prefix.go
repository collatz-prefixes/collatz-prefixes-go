package prefix

import (
	c "collatz-prefixes/internal/common"
	"math/big"
)

// Returns the prefix of two numbers
func PrefixFind(n *big.Int, m *big.Int) []uint {
	ans := make([]uint, 0)
	twos := uint(0)
	for {
		if c.IsEven(n) && c.IsEven(m) {
			// both are even
			twos++
			n.Rsh(n, 1)
			m.Rsh(m, 1)
		} else if !c.IsEven(n) && !c.IsEven(m) {
			// both are odd
			ans = append(ans, twos)
			n.Mul(n, c.THREE).Add(n, c.ONE)
			m.Mul(m, c.THREE).Add(m, c.ONE)
		} else {
			// they are of different parity
			break
		}
	}
	return ans
}

func PrefixIterate(n *big.Int, pf []uint) *big.Int {
	if len(pf) == 0 {
		return n
	}
	n.Div(n, c.T.Rsh(c.ONE, pf[0]))
	for i := 1; i < len(pf); i++ {
		n.Mul(n, c.THREE).Add(n, c.ONE)
		n.Div(n, c.T.Rsh(c.ONE, pf[i]-pf[i-1]))
	}
	return n
}

func PrefixMapToNum(pf []uint) *big.Int {
	ans := big.NewInt(0)
	for i := 0; i < len(pf); i++ {
		ans.Add(ans, c.T.Rsh(c.ONE, pf[i]))
	}
	return ans
}

func PrefixMapFromNum(n *big.Int) []uint {
	ans := make([]uint, 0)
	var b uint = 0 // TODO: better var name?
	for n.Cmp(c.ZERO) == 1 {
		if !c.IsEven(n) {
			ans = append(ans, b)
		}
		b++
		n.Rsh(n, 1)
	}
	return ans
}

func PrefixAdd(pf1 []uint, pf2 []uint) []uint {
	if len(pf1) == 0 {
		return pf2
	}
	pf1last := pf1[len(pf1)-1]
	pf1[len(pf1)-1] += pf2[0]
	for i := 1; i < len(pf2); i++ {
		pf1 = append(pf1, pf2[i]+pf1last)
	}
	return pf1
}
