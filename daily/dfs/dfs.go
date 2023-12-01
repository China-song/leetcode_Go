package dfs

import "math"

/*
SmallestMissingValueSubtree2003
2003. Smallest Missing Genetic Value in Each Subtree
*/

func SmallestMissingValueSubtree2003(parents []int, nums []int) []int {
	// 0的所有值包括自身+其children的所有值
	//      1           2 3 4
	// 递归求解 并保存

	// 构造children []int
	n := len(parents)
	children := make([][]int, n)
	// children[i] = {} 表示无子节点
	for i := 1; i < n; i++ {
		// i is child of parents[i]
		children[parents[i]] = append(children[parents[i]], i)
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}

	var dfs func(int) (map[int]bool, int) // return geneSet of node and ans[node]
	dfs = func(node int) (map[int]bool, int) {
		geneSet := map[int]bool{nums[node]: true}

		for _, child := range children[node] {
			childGeneSet, y := dfs(child)
			res[node] = max(res[node], y) // res[node]>=max(res[child])
			// 合并 将小集合 并到 大集合中
			if len(childGeneSet) > len(geneSet) {
				geneSet, childGeneSet = childGeneSet, geneSet
			}
			for gene, _ := range childGeneSet {
				geneSet[gene] = true
			}
		}
		for geneSet[res[node]] {
			res[node]++
		}
		return geneSet, res[node]
	}

	dfs(0)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
2304. Minimum Path Cost in a Grid
*/
func minPathCost(grid [][]int, moveCost [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	f[m-1] = grid[m-1]
	for i := m - 2; i >= 0; i-- {
		for j, g := range grid[i] {
			f[i][j] = math.MaxInt
			for k, c := range moveCost[g] {
				f[i][j] = min(f[i][j], grid[i][j]+c+f[i+1][k])
			}
		}
	}
	res := math.MaxInt
	for j := range f[0] {
		res = min(res, f[0][j])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
