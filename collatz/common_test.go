package collatz

import (
	"math/big"
	"testing"
)

func TestCommon(t *testing.T) {
	// parity
	if !IsEven(big.NewInt(0)) {
		t.Errorf("0 should be even.")
	}
	if !IsEven(big.NewInt(42)) {
		t.Errorf("42 should be even.")
	}
	if IsEven(big.NewInt(69)) {
		t.Errorf("69 should be odd.")
	}

	// is one
	if !IsOne(big.NewInt(1)) {
		t.Errorf("1 should be 1.")
	}
	if IsOne(big.NewInt(42)) {
		t.Errorf("42 is not 1.")
	}
}
