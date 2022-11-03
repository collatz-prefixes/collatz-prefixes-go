package main

import (
	collatz "collatz-prefixes/pkg/collatz"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(collatz.CollatzECF(big.NewInt(3)))
	fmt.Println(collatz.CollatzECFtoN(collatz.CollatzECF(big.NewInt(3))))
}
