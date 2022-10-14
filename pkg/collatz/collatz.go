package collatz

import (
	c "collatz-prefixes/internal/common"
	"math/big"
)

func CollatzLength(n *big.Int) uint {
	var ans uint = 0
	for n.Cmp(c.ONE) != 0 {
		ans++
		if c.IsEven(n) {
			n.Rsh(n, 1)
		} else {
			n.Mul(n, c.THREE).Add(n, c.ONE)
		}
	}
	return ans
}

func CollatzSeqeunce(n *big.Int) []*big.Int {
	ans := make([]*big.Int, 0)
	for n.Cmp(c.ONE) != 0 {
		ans = append(ans, c.Copy(n))
		if c.IsEven(n) {
			n.Rsh(n, 1)
		} else {
			n.Mul(n, c.THREE).Add(n, c.ONE)
		}
	}
	return ans
}

func CollatzReducedSeqeunce(n *big.Int) []*big.Int {
	ans := make([]*big.Int, 0)
	for n.Cmp(c.ONE) != 0 {
		if c.IsEven(n) {
			n.Rsh(n, 1)
		} else {
			ans = append(ans, c.Copy(n))
			n.Mul(n, c.THREE).Add(n, c.ONE)
		}
	}
	ans = append(ans, c.Copy(c.ONE))
	return ans
}

func CollatzECF(n *big.Int) []uint {
	ans := make([]uint, 0)
	twos := uint(0)
	for n.Cmp(c.ONE) != 0 {
		if c.IsEven(n) {
			twos++
			n.Rsh(n, 1)
		} else {
			ans = append(ans, twos)
			n.Mul(n, c.THREE).Add(n, c.ONE)
		}
	}
	ans = append(ans, twos)
	return ans
}

func CollatzECFtoN(ecf []uint) *big.Int {
	ans := big.NewInt(1)
	for i := len(ecf) - 1; i > 0; i-- {
		ans.Lsh(ans, ecf[i]-ecf[i-1])
		ans.Sub(ans, c.ONE).Div(ans, c.THREE)
	}
	ans.Lsh(ans, ecf[0])
	return ans
}
