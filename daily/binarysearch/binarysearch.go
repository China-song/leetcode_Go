package binarysearch

import "sort"

/*
HIndex275
275. H-Index II
*/
func HIndex275(citations []int) int {
	n := len(citations)
	left, right := 0, len(citations)
	var mid int
	for left < right {
		mid = (left + right + 1) >> 1
		if citations[n-mid] >= mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func successfulPairs(spells []int, potions []int, success int64) (pairs []int) {
	/* 方法一: 二分查找
	sort.Ints(potions)
	res := make([]int, len(spells))
	for i, x := range spells {
		res[i] = len(potions) - sort.SearchInts(potions, int((success-1)/int64(x)+1))
	}
	return res
	*/

	// 方法二: 双指针
	res := make([]int, len(spells))
	idx := make([]int, len(spells))
	for i := 0; i < len(spells); i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return spells[idx[i]] < spells[idx[j]]
	}) // sort idx in ascending order
	sort.Slice(potions, func(i, j int) bool {
		return potions[i] > potions[j]
	}) // sort potions in descending order
	j := 0
	for _, p := range idx {
		v := spells[p]
		for j < len(potions) && int64(v*potions[j]) >= success {
			j++
		}
		res[p] = j
	}
	return res
}

/*
34. Find First and Last Position of Element in Sorted Array
*/

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1 // nums[right+1] >= target
		} else {
			left = mid + 1 // nums[left-1] < target
		}
	}
	// left = right+1   right left
	return left
}

func searchRange(nums []int, target int) []int {
	leftIdx := binarySearch(nums, target)
	if leftIdx == len(nums) || nums[leftIdx] != target {
		return []int{-1, -1}
	}
	rightIdx := binarySearch(nums, target+1) - 1
	return []int{leftIdx, rightIdx}
}
