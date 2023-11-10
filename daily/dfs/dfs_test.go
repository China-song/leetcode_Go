package dfs

import "testing"

func TestSmallestMissingValueSubtree2003(t *testing.T) {
	cases := struct {
		parents []int
		nums    []int
	}{
		parents: []int{-1, 0, 0, 2},
		nums:    []int{1, 2, 3, 4},
	}

	SmallestMissingValueSubtree2003(cases.parents, cases.nums)
}
