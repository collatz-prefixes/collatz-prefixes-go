package prefix

import (
	"collatzprefixes/internal/equals"
	"collatzprefixes/pkg/collatz"
	"math/big"
	"testing"
)

// Find Prefix by simply comparing the ECFs
func prefixBrute(a []uint, b []uint) []uint {
	// get the minimum length
	l := len(a)
	if len(b) < l {
		l = len(b)
	}

	ans := make([]uint, 0)
	for i := 0; i < l; i++ {
		if a[i] == b[i] {
			ans = append(ans, a[i])
		} else {
			break
		}
	}

	return ans
}

func TestPrefix(t *testing.T) {
	cases := []struct {
		n *big.Int
		m *big.Int
	}{
		{big.NewInt(1), big.NewInt(2)},
		{big.NewInt(3), big.NewInt(12)},
		{big.NewInt(8), big.NewInt(16)},
		{big.NewInt(27), big.NewInt(37)},
	}

	for _, test := range cases {
		// expected prefix found via ECFs
		pf := prefixBrute(collatz.ECF(test.n), collatz.ECF(test.m))

		if !equals.Uints(pf, Find(test.n, test.m)) {
			t.Errorf("Prefix is wrong.")
		}

		if !equals.Uints(pf, Find(test.m, test.n)) {
			t.Errorf("PrefixFind should be commutative.")
		}

		// test both numbers for ECF iteration
		if Iterate(test.n, collatz.ECF(test.n)).Cmp(big.NewInt(1)) != 0 {
			t.Errorf("Iterating over ECF should result in 1.")
		}
		if Iterate(test.m, collatz.ECF(test.m)).Cmp(big.NewInt(1)) != 0 {
			t.Errorf("Iterating over ECF should result in 1.")
		}
	}
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
		{[]uint{4}, []uint{2, 4}, []uint{6, 8}},
		{[]uint{0, 1, 5}, []uint{0, 1, 3}, []uint{0, 1, 5, 6, 8}},
	}

	for _, test := range cases {
		if !equals.Uints(Add(test.p1, test.p2), test.result) {
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
		if ToNum(test.pf).Cmp(test.k) != 0 {
			t.Errorf("Wrong number from prefix.")
		}

		if !equals.Uints(FromNum(test.k), test.pf) {
			t.Errorf("Wrong prefix from number.")
		}
	}
}
