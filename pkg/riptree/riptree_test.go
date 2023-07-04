package riptree

import (
	"collatzprefixes/internal/common"
	"collatzprefixes/internal/equals"
	"collatzprefixes/pkg/prefix"
	"collatzprefixes/pkg/utils"
	"math/big"
	"testing"
)

func TestRiptree(t *testing.T) {
	cases := []struct {
		n  *big.Int
		pf []uint
	}{
		// edge
		{big.NewInt(1), []uint{0}},
		// power of 2
		{big.NewInt(2), []uint{1}},
		{big.NewInt(8), []uint{3}},
		// others
		{big.NewInt(5), []uint{0}},
		{big.NewInt(3), []uint{0, 1}},
		{big.NewInt(7), []uint{0, 1, 2}},
		{big.NewInt(27), []uint{0, 1, 3, 4}},
		{big.NewInt(321), []uint{0, 2, 4}},
		{big.NewInt(322), []uint{1, 3, 5, 6, 8}},
	}

	for _, test := range cases {
		pf := PrefixFind(test.n, utils.NTOP(test.n))

		if !equals.Uints(pf, test.pf) {
			t.Errorf("Wrong prefix.")
		}

		if common.IsEven(prefix.Iterate(test.n, pf)) {
			t.Errorf("Result of prefix iteration should be odd.")
		}

	}
}

func TestNextInPath(t *testing.T) {
	cases := []struct {
		n *big.Int
		k *big.Int
	}{
		{big.NewInt(1), big.NewInt(2)},
		{big.NewInt(2), big.NewInt(4)},
		{big.NewInt(3), big.NewInt(7)},
		{big.NewInt(7), big.NewInt(15)},
		{big.NewInt(5), big.NewInt(13)},
	}

	for _, test := range cases {
		if NextInPath(test.n, utils.NTOP(test.n)).Cmp(test.k) != 0 {
			t.Errorf("Wrong number for next in path.")
		}
	}
}
