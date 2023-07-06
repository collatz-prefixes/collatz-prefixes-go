package collatz

import (
	"collatzprefixes/internal/common"
	"collatzprefixes/internal/equals"
	"math/big"
	"testing"
)

func TestCollatzECF(t *testing.T) {
	cases := []struct {
		n   *big.Int
		ecf []uint
	}{
		// edge
		{big.NewInt(1), []uint{0}},
		// power of two
		{big.NewInt(16), []uint{4}},
		// odd number
		{big.NewInt(3), []uint{0, 1, 5}},
		// even number
		{big.NewInt(12), []uint{2, 3, 7}},
		// large sequence https://oeis.org/A008884
		{big.NewInt(27), []uint{0, 1, 3, 4, 5, 6, 7, 9, 11, 12, 14, 15, 16, 18, 19, 20, 21, 23, 26, 27, 28, 30, 31, 33, 34, 35, 36, 37, 38, 41, 42, 43, 44, 48, 50, 52, 56, 59, 60, 61, 66, 70}},
	}

	for _, test := range cases {
		ecf := ECF(test.n)
		if !equals.Uints(ecf, test.ecf) {
			t.Errorf("Wrong ECF from number.\t%x != %x", test.ecf, ecf)
		}

		n := ECFtoN(test.ecf)
		if n.Cmp(test.n) != 0 {
			t.Errorf("Wrong number from ECF.\t%d != %d", test.n, n)
		}
	}

}

func TestCollatzSequences(t *testing.T) {
	cases := []struct {
		n    *big.Int
		seq  []*big.Int
		rseq []*big.Int
	}{
		// edge
		{big.NewInt(1), common.ToBigInts([]int64{1}), common.ToBigInts([]int64{1})},
		// power of two
		{big.NewInt(8), common.ToBigInts([]int64{8, 4, 2, 1}), common.ToBigInts([]int64{8, 1})},
		// odd number
		{big.NewInt(5), common.ToBigInts([]int64{5, 16, 8, 4, 2, 1}), common.ToBigInts([]int64{5, 1})},
		// even number
		{big.NewInt(28), common.ToBigInts([]int64{28, 14, 7, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1}), common.ToBigInts([]int64{28, 7, 11, 17, 13, 5, 1})},
		// large sequence https://oeis.org/A008884
		{big.NewInt(27), common.ToBigInts([]int64{27, 82, 41, 124, 62, 31, 94, 47, 142, 71, 214, 107, 322, 161, 484, 242, 121, 364, 182, 91, 274, 137, 412, 206, 103, 310, 155, 466, 233, 700, 350, 175, 526, 263, 790, 395, 1186, 593, 1780, 890, 445, 1336, 668, 334, 167, 502, 251, 754, 377, 1132, 566, 283, 850, 425, 1276, 638, 319, 958, 479, 1438, 719, 2158, 1079, 3238, 1619, 4858, 2429, 7288, 3644, 1822, 911, 2734, 1367, 4102, 2051, 6154, 3077, 9232, 4616, 2308, 1154, 577, 1732, 866, 433, 1300, 650, 325, 976, 488, 244, 122, 61, 184, 92, 46, 23, 70, 35, 106, 53, 160, 80, 40, 20, 10, 5, 16, 8, 4, 2, 1}), common.ToBigInts([]int64{27, 41, 31, 47, 71, 107, 161, 121, 91, 137, 103, 155, 233, 175, 263, 395, 593, 445, 167, 251, 377, 283, 425, 319, 479, 719, 1079, 1619, 2429, 911, 1367, 2051, 3077, 577, 433, 325, 61, 23, 35, 53, 5, 1})},
	}

	for _, test := range cases {
		seq := Seqeunce(test.n)
		if !equals.Bigints(seq, test.seq) {
			t.Errorf("Wrong sequence.\t%x != %x", test.seq, seq)
		}

		rseq := ReducedSeqeunce(test.n)
		if !equals.Bigints(rseq, test.rseq) {
			t.Errorf("Wrong reduced sequence.\t%x != %x", test.rseq, rseq)
		}

		test_l := len(test.seq) - 1 // last is ignored
		l := Length(test.n)
		if l != test_l {
			t.Errorf("Wrong length.\t%d != %d", test_l, l)
		}
	}

}
