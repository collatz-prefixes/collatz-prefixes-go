package collatz

import (
	"math/big"
	"testing"
)

func TestPrefix(t *testing.T) {

}

func TestPrefixAdd(t *testing.T) {
	cases := []struct {
		p1     []uint
		p2     []uint
		result []uint
	}{
		// edge
		{[]uint{}, []uint{1}, []uint{1}},
		{[]uint{1}, []uint{}, []uint{1}},
		{nil, []uint{1}, []uint{1}},
		{[]uint{1}, nil, []uint{1}},
		// normal cases
		{[]uint{2, 4}, []uint{4}, []uint{2, 8}},
		{[]uint{0, 1, 5}, []uint{0, 1, 3}, []uint{0, 1, 5, 6, 8}},
	}

	for _, test := range cases {
		if !EqualUints(PrefixAdd(test.p1, test.p2), test.result) {
			t.Errorf("Wrong result.")
		}
	}

}

func TestPrefixMap(t *testing.T) {
	cases := []struct {
		k  *big.Int
		pf []uint
	}{
		// edge
		{big.NewInt(1), []uint{0}},
		// power of 2
		{big.NewInt(16), []uint{4}},
		// odds
		{big.NewInt(27), []uint{0, 1, 3, 4}},
		{big.NewInt(35), []uint{0, 1, 5}},
		// evens
		{big.NewInt(12), []uint{2, 3}},
		{big.NewInt(190), []uint{1, 2, 3, 4, 5, 7}},
	}

	for _, test := range cases {
		if PrefixMapToNum(test.pf).Cmp(test.k) != 0 {
			t.Errorf("Wrong number from prefix.")
		}

		if !EqualUints(PrefixMapFromNum(test.k), test.pf) {
			t.Errorf("Wrong prefix from number.")
		}
	}
}
