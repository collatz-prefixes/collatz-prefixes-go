package equals

import (
	"math/big"
	"testing"
)

func TestEquals(t *testing.T) {
	if !Uints([]uint{1, 2, 3}, []uint{1, 2, 3}) {
		t.Errorf("Should have been equal.")
	}
	if Uints([]uint{1, 2, 3}, []uint{7, 8, 9}) {
		t.Errorf("Should NOT have been equal.")
	}

	if !Bools([]bool{true, true}, []bool{true, true}) {
		t.Errorf("Should have been equal.")
	}
	if Bools([]bool{false, true}, []bool{true, false}) {
		t.Errorf("Should NOT have been equal.")
	}

	if !Bigints([]*big.Int{big.NewInt(1), big.NewInt(2)}, []*big.Int{big.NewInt(1), big.NewInt(2)}) {
		t.Errorf("Should have been equal.")
	}
	if Bigints([]*big.Int{big.NewInt(1), big.NewInt(2)}, []*big.Int{big.NewInt(1), big.NewInt(9)}) {
		t.Errorf("Should NOT have been equal.")
	}
}
