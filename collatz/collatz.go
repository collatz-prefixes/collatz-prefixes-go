package collatz

import "math/big"

// Collatz length is the number of iterations it takes to reach n to 1,
func CollatzLength(n *big.Int) uint {
	n = Copy(n)
	var ans uint = 0
	for !IsOne(n) {
		ans++
		if IsEven(n) {
			n.Rsh(n, 1)
		} else {
			n.Mul(n, THREE).Add(n, ONE)
		}
	}
	return ans
}

// Collatz Sequence is the array of numbers seen during iterations until 1 is reached.
func CollatzSeqeunce(n *big.Int) []*big.Int {
	n = Copy(n)
	ans := make([]*big.Int, 0)
	for !IsOne(n) {
		ans = append(ans, Copy(n))
		if IsEven(n) {
			n.Rsh(n, 1)
		} else {
			n.Mul(n, THREE).Add(n, ONE)
		}
	}
	return ans
}

// Reduced Collatz Sequence is the array of odd numbers seen during iterations until 1 is reached.
func CollatzReducedSeqeunce(n *big.Int) []*big.Int {
	n = Copy(n)
	ans := make([]*big.Int, 0)
	if IsEven(n) {
		ans = append(ans, Copy(n))
	}

	for !IsOne(n) {
		if IsEven(n) {
			n.Rsh(n, 1)
		} else {
			ans = append(ans, Copy(n))
			n.Mul(n, THREE).Add(n, ONE)
		}
	}

	ans = append(ans, Copy(ONE))
	return ans
}

// Find ECF (Exponential Canonical Form) of a number.
func CollatzECF(n *big.Int) []uint {
	n = Copy(n)
	ans := make([]uint, 0)
	twos := uint(0)
	for !IsOne(n) {
		if IsEven(n) {
			twos++
			n.Rsh(n, 1)
		} else {
			ans = append(ans, twos)
			n.Mul(n, THREE).Add(n, ONE)
		}
	}
	ans = append(ans, twos)
	return ans
}

// Compute a number from it's ECF.
func CollatzECFtoN(ecf []uint) *big.Int {
	ans := big.NewInt(1)
	for i := len(ecf) - 1; i > 0; i-- {
		ans.Lsh(ans, ecf[i]-ecf[i-1])
		ans.Sub(ans, ONE).Div(ans, THREE)
	}
	ans.Lsh(ans, ecf[0])
	return ans
}
