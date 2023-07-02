package collatz

import "math/big"

// Finds the nature of a path.
func PiptreeFindNature(p []bool, pf []uint, rpf uint) bool {
	n := PTON(p)
	iter_res := PrefixIterate(n, append(pf, rpf+1))
	return IsEven(iter_res)
}

// Finds the path from root to the node indexed by p in PIPTree, with the path length of the root node being equal to |p|.
//
// It starts from the target, and in a loop either does `m/2` or `(m-1)/2` until it reaches 1.
// This gives the path from that number to root, so we reverse that to obtain the road from path to target.
func PiptreeGetRootDirections(p []bool) []bool {
	ans := make([]bool, 0)
	for i := BTON(p); i.Cmp(ONE) == 1; {
		if IsEven(i) {
			i.Rsh(i, 1)
			ans = append(ans, false) // left
		} else {
			i.Rsh(i.Sub(i, ONE), 1)
			ans = append(ans, true) // right
		}
	}
	REVERSE(ans)
	return ans

}

// Finds the prefix of a number using PIPTree properties.
func PiptreePrefixFind(n *big.Int, p []bool) []uint {
	// TODO: does this always work for even larger paths?
	if PTON(p).Cmp(n) != 0 {
		panic("Path does not match the number")
	}

	n = Copy(n)
	if ISPOW2(n) {
		ans := uint(0)
		for ; n.Cmp(ONE) == 1; ans++ {
			n.Rsh(n, 1)
		}
		return []uint{ans}
	} else {
		// directions from root to p
		dir := PiptreeGetRootDirections(p)

		// root prefix
		rpf := uint(len(p) - 1)

		// root number
		r := new(big.Int).Rsh(ONE, rpf)

		// root path
		rp := make([]bool, len(p)) // values false by default
		rp[len(p)-1] = true

		// start from the root and work your way to the target
		cur_n := r
		cur_p := rp
		cur_pf := []uint{rpf}

		for _, d := range dir {
			// nature of the current node
			nat := PiptreeFindNature(cur_p, cur_pf, rpf)

			// decrement everything in the prefix
			for i := range cur_pf {
				cur_pf[i]--
			}

			// append root prefix if
			// BAD and RIGHT, or
			// GOOD and LEFT
			if (d && !nat) || (!d && nat) {
				cur_pf = append(cur_pf, rpf)
			}

			// div by 2
			cur_n.Rsh(cur_n, 1)

			// if GOOD, add root too
			if !d {
				cur_n.Add(cur_n, r)
			}

			// go to the next child
			i := 0
			for ; i < len(cur_p)-1; i++ {
				cur_p[i] = cur_p[i+1]
			}
			cur_p[i] = d
		}

		return cur_pf
	}

}
