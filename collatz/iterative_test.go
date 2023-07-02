package collatz

import (
	"math/big"
	"testing"
)

func TestIterativePrefix(t *testing.T) {
	cases := []struct {
		n *big.Int
	}{
		{big.NewInt(1)},
		{big.NewInt(5)},
		{big.NewInt(8)},
		{big.NewInt(27)},
		{big.NewInt(38)},
		{big.NewInt(186438726873)},
	}

	for _, test := range cases {
		test_ecf := CollatzECF(test.n)

		ecf_pf_rip := IterativePrefix(test.n, RiptreePrefixFind)
		if !EqualUints(ecf_pf_rip, test_ecf) {
			t.Errorf("ECF mismatch using RIPTree.\t%x != %x", test_ecf, ecf_pf_rip)
		}

		ecf_pf_pip := IterativePrefix(test.n, PiptreePrefixFind)
		if !EqualUints(ecf_pf_pip, test_ecf) {
			t.Errorf("ECF mismatch using PIPTree.\t%x != %x", test_ecf, ecf_pf_pip)
		}
	}
}

func TestIterativePath(t *testing.T) {
	cases := []struct {
		n *big.Int
	}{
		{big.NewInt(1)},
		{big.NewInt(5)},
		{big.NewInt(8)},
		{big.NewInt(27)},
		{big.NewInt(38)},
		{big.NewInt(186438726873)},
	}

	for _, test := range cases {
		test_ecf := CollatzECF(test.n)

		ecf_p_rip := IterativePathExtension(test.n, RiptreePrefixFind)
		if !EqualUints(ecf_p_rip, test_ecf) {
			t.Errorf("ECF mismatch using RIPTree.\t%x != %x", test_ecf, ecf_p_rip)
		}

		ecf_p_pip := IterativePathExtension(test.n, PiptreePrefixFind)
		if !EqualUints(ecf_p_pip, test_ecf) {
			t.Errorf("ECF mismatch using PIPTree.\t%x != %x", test_ecf, ecf_p_pip)
		}
	}
}
