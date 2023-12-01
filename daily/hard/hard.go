package hard

import "sort"

/*
2736. Maximum Sum Queries
*/
func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	type pair struct {
		x, y int
	}
	a := make([]pair, len(nums1))
	for i, x := range nums1 {
		a[i] = pair{x, nums2[i]}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].x > a[j].x
	})

	qid := make([]int, len(queries))
	for i := range qid {
		qid[i] = i
	}
	sort.Slice(qid, func(i, j int) bool {
		return queries[i][0] > queries[j][0]
	})

	ans := make([]int, len(queries)) // ans[i] is answer of queries[i]
	type data struct {
		y, s int
	}
	st := []data{} // 单调栈
	j := 0
	for _, i := range qid {
		// queries[i]: {xi, yi} xi >= other
		x, y := queries[i][0], queries[i][1]
		// 将所有nums1[j] >= xi 的 pair 记录下 用于 进一步搜索
		for ; j < len(a) && a[j].x >= x; j++ {
			// st 栈底到栈顶 y单调递增
			for len(st) > 0 && (a[j].y > st[len(st)-1].y && st[len(st)-1].s <= a[j].x+a[j].y) {
				st = st[:len(st)-1]
			}
			if len(st) == 0 || a[j].y > st[len(st)-1].y { // 要记录的pair的y需要比之前记录的更大，不满足的不可能是答案
				st = append(st, data{a[j].y, a[j].x + a[j].y}) // st 栈底到栈顶 y单调递增
			}
		}
		p := sort.Search(len(st), func(i int) bool {
			return st[i].y >= y
		})
		if p < len(st) {
			ans[i] = st[p].s
		} else {
			ans[i] = -1
		}
	}
	return ans
}
