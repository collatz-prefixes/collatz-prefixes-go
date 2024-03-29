package common

import "math/big"

var (
	// shorthand for big.NewInt(0)
	ZERO = big.NewInt(0)
	// shorthand for big.NewInt(1)
	ONE = big.NewInt(1)
	// shorthand for big.NewInt(2)
	TWO = big.NewInt(2)
	// shorthand for big.NewInt(3)
	THREE = big.NewInt(3)
)

// Copy the pointed big int, returns a new pointer with the same value.
func Copy(n *big.Int) *big.Int {
	return new(big.Int).Set(n)
}

// Returns true if given number is even.
func IsEven(n *big.Int) bool {
	return n.Bit(0) == 0
}

// Returns true if given number is equal to 1.
func IsOne(n *big.Int) bool {
	return n.Cmp(ONE) == 0
}

// Maps an array of `int64`as to an array of bigints.
func ToBigInts(arr []int64) []*big.Int {
	ans := make([]*big.Int, len(arr))
	for i, v := range arr {
		ans[i] = big.NewInt(v)
	}
	return ans
}
