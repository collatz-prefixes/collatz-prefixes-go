package equals

import "math/big"

// Compare two arrays of uints.
func Uints(a, b []uint) bool {
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

// Compare two arrays of bigints.
func Bigints(a, b []*big.Int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.Cmp(b[i]) != 0 {
			return false
		}
	}
	return true
}

// Compare two arrays of booleans.
func Bools(a, b []bool) bool {
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
