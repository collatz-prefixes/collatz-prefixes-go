package utils

import (
	"collatzprefixes/internal/equals"
	"math/big"
	"testing"
)

func TestPath(t *testing.T) {
	cases := []struct {
		n *big.Int
		p []bool
	}{
		// edge
		{big.NewInt(1), []bool{}},
		// power of 2
		{big.NewInt(2), []bool{false}},
		{big.NewInt(4), []bool{false, false}},
		// odds
		{big.NewInt(3), []bool{true, false}},
		{big.NewInt(5), []bool{true, true, false}},
		{big.NewInt(7), []bool{true, false, false}},
		{big.NewInt(9), []bool{true, true, true, false}},
		// evens
		{big.NewInt(6), []bool{false, true, false}},
		{big.NewInt(10), []bool{false, true, true, false}},
		{big.NewInt(12), []bool{false, false, true, false}},
	}

	for _, test := range cases {
		if !equals.Bools(NTOP(test.n), test.p) {
			t.Errorf("Wrong path from number.")
		}
		if PTON(test.p).Cmp(test.n) != 0 {
			t.Errorf("Wrong number from path.")
		}
	}
}

func TestBinary(t *testing.T) {
	cases := []struct {
		n *big.Int
		b []bool
	}{
		{big.NewInt(1), []bool{true}},
		{big.NewInt(3), []bool{true, true}},
		{big.NewInt(8), []bool{true, false, false, false}},
		{big.NewInt(15), []bool{true, true, true, true}},
	}

	for _, test := range cases {
		if !equals.Bools(NTOB(test.n), test.b) {
			t.Errorf("Wrong binary from number.")
		}
		if BTON(test.b).Cmp(test.n) != 0 {
			t.Errorf("Wrong number from binary.")
		}
	}
}

func TestPow2(t *testing.T) {
	cases := []struct {
		yes *big.Int
		no  *big.Int
	}{
		{big.NewInt(1), big.NewInt(3)},
		{big.NewInt(2), big.NewInt(5)},
		{big.NewInt(4), big.NewInt(7)},
		{big.NewInt(16), big.NewInt(19)},
	}

	for _, test := range cases {
		if !ISPOW2(test.yes) {
			t.Errorf("ISPOW2 failed.")
		}
		if ISPOW2(test.no) {
			t.Errorf("ISPOW2 failed.")
		}
	}
}

func TestBools(t *testing.T) {
	if !equals.Bools(FLIP([]bool{false, false}), []bool{true, true}) {
		t.Errorf("FLIP failed.")
	}
	if !equals.Bools(REVERSE([]bool{false, true}), []bool{true, false}) {
		t.Errorf("REVERSE failed.")
	}
}
