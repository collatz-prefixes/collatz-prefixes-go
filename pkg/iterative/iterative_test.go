package iterative

import (
	"collatzprefixes/internal/equals"
	"collatzprefixes/pkg/collatz"
	"collatzprefixes/pkg/piptree"
	"collatzprefixes/pkg/riptree"
	"math/big"
	"testing"
)

func TestIteratives(t *testing.T) {
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

		test_ecf := collatz.ECF(test.n)

		ecf_pf_rip := Prefix(test.n, riptree.PrefixFind)
		if !equals.Uints(ecf_pf_rip, test_ecf) {
			t.Errorf("ECF mismatch using Prefix + RIPTree.\t%x != %x", test_ecf, ecf_pf_rip)
		}

		ecf_pf_pip := Prefix(test.n, piptree.PrefixFind)
		if !equals.Uints(ecf_pf_pip, test_ecf) {
			t.Errorf("ECF mismatch using Prefix + PIPTree.\t%x != %x", test_ecf, ecf_pf_pip)
		}

		ecf_p_rip := PathExtension(test.n, riptree.PrefixFind)
		if !equals.Uints(ecf_p_rip, test_ecf) {
			t.Errorf("ECF mismatch using Path + RIPTree.\t%x != %x", test_ecf, ecf_p_rip)
		}

		ecf_p_pip := PathExtension(test.n, piptree.PrefixFind)
		if !equals.Uints(ecf_p_pip, test_ecf) {
			t.Errorf("ECF mismatch using Path + PIPTree.\t%x != %x", test_ecf, ecf_p_pip)
		}
	}
}
