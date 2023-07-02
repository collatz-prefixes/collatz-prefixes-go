package collatz

import (
	"math/big"
	"testing"
)

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

func TestCollatzECF(t *testing.T) {
	type ecfTest struct {
		n   *big.Int
		ecf []uint
	}

	var ecfTests = []ecfTest{
		{big.NewInt(3), []uint{0, 1, 5}},
		{big.NewInt(5), []uint{0, 4}},
		{big.NewInt(6), []uint{1, 2, 6}},
		{big.NewInt(16), []uint{4}},
	}

	for _, test := range ecfTests {
		if !equal(CollatzECF(test.n), test.ecf) {
			t.Errorf("Wrong ECF for this number.")
		}

		if CollatzECFtoN(test.ecf).Cmp(test.n) != 0 {
			t.Errorf("Wrong number for this ECF.")
		}
	}

}
