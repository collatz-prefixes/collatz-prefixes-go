package riptree

import (
	c "collatz-prefixes/internal/common"
	"math/big"
)

func RiptreeNextInPath(n *big.Int, p []bool) *big.Int {
	return n.Add(n, c.T.Exp(c.TWO, big.NewInt(int64(len(p))), nil))
}

// TODO
