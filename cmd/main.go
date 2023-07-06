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

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Expected two arguments: function & number")
		return
	}

	function := args[0]
	n, ok := new(big.Int).SetString(args[1], 10)
	if !ok {
		fmt.Println("SetString failed")
		return
	}

	switch function {
	case "len":
		fmt.Println(collatz.Length(n))
	case "seq":
		fmt.Println(collatz.Seqeunce(n))
	case "rdseq":
		fmt.Println(collatz.ReducedSeqeunce(n))
	case "ecf":
		fmt.Println(collatz.ECF(n))
	case "path":
		fmt.Println(utils.ToPath(n))
	case "map":
		fmt.Println(prefix.FromNum(n))
	case "pf-map":
		pf := riptree.PrefixFind(n, utils.ToPath(n))
		fmt.Println(prefix.ToNum(pf))
	case "pf-rip":
		fmt.Println(riptree.PrefixFind(n, utils.ToPath(n)))
	case "pf-pip":
		fmt.Println(piptree.PrefixFind(n, utils.ToPath(n)))
	case "ecf-pf-rip":
		fmt.Println(iterative.Prefix(n, riptree.PrefixFind))
	case "ecf-pf-pip":
		fmt.Println(iterative.Prefix(n, piptree.PrefixFind))
	case "ecf-path-rip":
		fmt.Println(iterative.PathExtension(n, piptree.PrefixFind))
	case "ecf-path-pip":
		fmt.Println(iterative.PathExtension(n, piptree.PrefixFind))
	default:
		fmt.Println("Unknown function.")
	}

}
