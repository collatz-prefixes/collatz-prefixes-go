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
	} else if function == "rseq" {
		fmt.Println(collatz.CollatzReducedSeqeunce(n))
	} else if function == "path" {
		fmt.Println(collatz.NTOP(n))
	} else if function == "pf" {
		m, ok := new(big.Int).SetString(args[1], 10)
		if !ok {
			fmt.Println("SetString failed")
			return
		}
		fmt.Println(collatz.PrefixFind(n, m))
	} else if function == "mappf" {
		fmt.Println(collatz.PrefixMapFromNum(n))
	} else if function == "rippf" {
		fmt.Println(collatz.RiptreePrefixFind(n, collatz.NTOP(n)))
	} else if function == "pippf" {
		fmt.Println(collatz.PiptreePrefixFind(n, collatz.NTOP(n)))
	} else if function == "iter" {
		fmt.Println(collatz.IterativePathExtension(n, collatz.PiptreePrefixFind))
	} else {
		fmt.Println("Unknown function.")
	}

}
