package middle

import "sort"

/*
SumDistance2731
2731. Movement of Robots
*/
func SumDistance2731(nums []int, s string, d int) int {
	const mod int = 1e9 + 7
	n := len(nums)
	// calculate coordinates of robots after d seconds
	pos := make([]int, n)
	for i, ch := range s {
		if ch == 'L' {
			pos[i] = nums[i] - d
		} else {
			pos[i] = nums[i] + d
		}
	}

	// sort pos in increasing order: pos[0] and pos[n-1] are the edge
	sort.Ints(pos)

	// calculate distance
	res := 0
	for i := 1; i < n; i++ {
		res += (pos[i] - pos[i-1]) * i % mod * (n - i) % mod
		res %= mod
	}

	return res
}
