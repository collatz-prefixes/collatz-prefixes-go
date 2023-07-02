package prefix

import (
	"collatzprefixes/internal/common"
	"collatzprefixes/pkg/collatz"
	"math/big"
)

// Returns the prefix of two numbers.
//
// The prefix can be thought of as common prefix of the ECFs.
// As an example, ECF(3) = [0, 1, 5] and ECF(7) = [0, 1, 2, 4, 7, 11].
// The common prefix here is [0, 1], thus prefix(3,7) = prefix(7,3) = [0, 1].
func Find(n *big.Int, m *big.Int) []uint {
	if n.Cmp(m) == 0 {
		return collatz.ECF(n)
	}

	ans := make([]uint, 0)
	twos := uint(0)
	n = common.Copy(n)
	m = common.Copy(m)
	for {
		if common.IsEven(n) && common.IsEven(m) {
			// both are even
			twos++
			n.Rsh(n, 1)
			m.Rsh(m, 1)
		} else if !common.IsEven(n) && !common.IsEven(m) {
			// both are odd
			ans = append(ans, twos)
			n.Mul(n, common.THREE).Add(n, common.ONE)
			m.Mul(m, common.THREE).Add(m, common.ONE)
		} else {
			// they are of different parity
			break
		}
	}
	return ans
}

// Iterates a number through a prefix.
//
// If the prefix is equal to ECF of the number, the result is expected to be 1.
func Iterate(n *big.Int, pf []uint) *big.Int {
	if len(pf) == 0 {
		return n
	}
	n = common.Copy(n)

	// R_0 function
	n.Div(n, new(big.Int).Lsh(common.ONE, pf[0]))

	// R function
	for i := 1; i < len(pf); i++ {
		n.Mul(n, common.THREE).Add(n, common.ONE)
		n.Div(n, new(big.Int).Lsh(common.ONE, pf[i]-pf[i-1]))
	}
	return n
}

// Bijective mapping from a list of ascending numbers to an integer
func ToNum(pf []uint) *big.Int {
	ans := big.NewInt(0)
	for i := 0; i < len(pf); i++ {
		ans.Add(ans, new(big.Int).Lsh(common.ONE, pf[i]))
	}
	return ans
}

// Bijective mapping from a list of ascending numbers to an integer
func FromNum(k *big.Int) []uint {
	k = common.Copy(k)
	ans := make([]uint, 0)
	for bitPos := uint(0); k.Cmp(common.ZERO) == 1; bitPos++ {
		if !common.IsEven(k) {
			ans = append(ans, bitPos)
		}
		k.Rsh(k, 1)
	}
	return ans
}

// Add two prefixes.
func Add(pf1 []uint, pf2 []uint) []uint {
	// edge cases
	if len(pf1) == 0 {
		return pf2
	}
	if len(pf2) == 0 {
		return pf1
	}

	ans := make([]uint, len(pf1))
	copy(ans, pf1)

	last := pf1[len(pf1)-1]
	ans[len(pf1)-1] += pf2[0]
	for i := 1; i < len(pf2); i++ {
		ans = append(ans, pf2[i]+last)
	}

	return ans
}
