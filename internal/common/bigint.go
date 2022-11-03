package common

import "math/big"

var (
	ZERO  = big.NewInt(0)
	ONE   = big.NewInt(1)
	TWO   = big.NewInt(2)
	THREE = big.NewInt(3)
	T     = big.NewInt(0) // tmp
)

// Copy the pointed big int, returns a new pointer
func Copy(n *big.Int) *big.Int {
	ans := big.NewInt(0)
	return ans.Set(n)
}

// Returns true if given number is even
func IsEven(n *big.Int) bool {
	return n.Bit(0) == 0
}
