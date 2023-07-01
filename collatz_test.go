package collatz_prefixes

import (
	"math/big"
	"testing"
)

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestCollatzLength(t *testing.T) {
	if !equal(CollatzECF(big.NewInt(3)), []uint{0, 1, 5}) {
		t.Fail()
	}

}
