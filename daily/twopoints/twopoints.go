package twopoints

import "sort"

/*
795. Number of Subarrays with Bounded Maximum
*/
func numSubarrayBoundedMax(nums []int, left int, right int) (res int) {
	// [l, r] 维护一个值maxNum: left <= maxNum <= right
	return

}

/*
2824. Count Pairs Whose Sum is Less than Target
*/
func countPairs(nums []int, target int) (ans int) {
	// -1 1 1 2 3
	sort.Ints(nums)
	for j := 1; j < len(nums); j++ {
		ans += sort.SearchInts(nums[:j], target-nums[j]) // 找nums[i] + nums[j] >= target 的最小i
		// [0..i-1]都是满足nums[i]+nums[j]<target的i
	}
	return ans
}
