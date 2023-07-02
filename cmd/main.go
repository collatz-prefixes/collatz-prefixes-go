package main

import (
	collatz "collatzprefix/collatz"
	"fmt"
	"math/big"
	"os"
)

func main() {
	args := os.Args[1:]

	// parse arguments
	function := args[0]
	n, ok := new(big.Int).SetString(args[1], 10)
	if !ok {
		fmt.Println("SetString failed")
		return
	}

	// process command
	if function == "ecf" {
		fmt.Println(collatz.CollatzECF(n))
	} else if function == "seq" {
		fmt.Println(collatz.CollatzSeqeunce(n))
	} else if function == "ecf-n" {
		fmt.Println(collatz.CollatzECFtoN([]uint{0, 1, 5}))
	} else if function == "rseq" {
		fmt.Println(collatz.CollatzReducedSeqeunce(n))
	} else if function == "path" {
		fmt.Println(collatz.NTOP(n))
	} else if function == "prefix" {
		m, ok := new(big.Int).SetString(args[1], 10)
		if !ok {
			fmt.Println("SetString failed")
			return
		}
		fmt.Println(collatz.PrefixFind(n, m))
	} else {
		fmt.Println("Unknown function.")
	}

}
