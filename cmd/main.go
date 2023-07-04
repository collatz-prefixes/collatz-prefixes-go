package main

import (
	"collatzprefixes/pkg/collatz"
	"collatzprefixes/pkg/iterative"
	"collatzprefixes/pkg/piptree"
	"collatzprefixes/pkg/prefix"
	"collatzprefixes/pkg/riptree"
	"collatzprefixes/pkg/utils"
	"fmt"
	"math/big"
	"os"
)

// TODO: use flag library to get arguments

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
		fmt.Println(collatz.ECF(n))
	} else if function == "seq" {
		fmt.Println(collatz.Seqeunce(n))
	} else if function == "rseq" {
		fmt.Println(collatz.ReducedSeqeunce(n))
	} else if function == "path" {
		fmt.Println(utils.NTOP(n))
	} else if function == "pf" {
		m, ok := new(big.Int).SetString(args[1], 10)
		if !ok {
			fmt.Println("SetString failed")
			return
		}
		fmt.Println(prefix.Find(n, m))
	} else if function == "mappf" {
		fmt.Println(prefix.FromNum(n))
	} else if function == "rippf" {
		fmt.Println(riptree.PrefixFind(n, utils.NTOP(n)))
	} else if function == "pippf" {
		fmt.Println(piptree.PrefixFind(n, utils.NTOP(n)))
	} else if function == "ripiter" {
		fmt.Println(iterative.Prefix(n, riptree.PrefixFind))
	} else if function == "pipiter" {
		fmt.Println(iterative.Prefix(n, piptree.PrefixFind))
	} else {
		fmt.Println("Unknown function.")
	}

}
