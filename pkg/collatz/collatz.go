package collatz

import (
	"collatzprefixes/internal/common"
	"math/big"
)

// Collatz length is the number of iterations it takes to reach n to 1.
func Length(n *big.Int) int {
	n = common.Copy(n)
	ans := 0
	for !common.IsOne(n) {
		ans++
		if common.IsEven(n) {
			n.Rsh(n, 1)
		} else {
			n.Mul(n, common.THREE).Add(n, common.ONE)
		}
	}
	return ans + 1
}

// Collatz Sequence is the array of numbers seen during iterations until 1 is reached.
func Seqeunce(n *big.Int) []*big.Int {
	n = common.Copy(n)
	ans := make([]*big.Int, 0)
	for !common.IsOne(n) {
		ans = append(ans, common.Copy(n))
		if common.IsEven(n) {
			n.Rsh(n, 1)
		} else {
			n.Mul(n, common.THREE).Add(n, common.ONE)
		}
	}
	ans = append(ans, common.Copy(n))
	return ans
}

// Reduced Collatz Sequence is the array of odd numbers seen during iterations until 1 is reached.
func ReducedSeqeunce(n *big.Int) []*big.Int {
	n = common.Copy(n)
	ans := make([]*big.Int, 0)
	if common.IsEven(n) {
		ans = append(ans, common.Copy(n))
	}

	for !common.IsOne(n) {
		if common.IsEven(n) {
			n.Rsh(n, 1)
		} else {
			ans = append(ans, common.Copy(n))
			n.Mul(n, common.THREE).Add(n, common.ONE)
		}
	}

	ans = append(ans, common.Copy(n))
	return ans
}

// Find ECF (Exponential Canonical Form) of a number.
func ECF(n *big.Int) []uint {
	n = common.Copy(n)
	ans := make([]uint, 0)
	twos := uint(0)
	for !common.IsOne(n) {
		if common.IsEven(n) {
			twos++
			n.Rsh(n, 1)
		} else {
			ans = append(ans, twos)
			n.Mul(n, common.THREE).Add(n, common.ONE)
		}
	}
	ans = append(ans, twos)
	return ans
}

// Compute a number from it's ECF.
func ECFtoN(ecf []uint) *big.Int {
	ans := big.NewInt(1)
	for i := len(ecf) - 1; i > 0; i-- {
		ans.Lsh(ans, ecf[i]-ecf[i-1])
		ans.Sub(ans, common.ONE).Div(ans, common.THREE)
	}
	ans.Lsh(ans, ecf[0])
	return ans
}
