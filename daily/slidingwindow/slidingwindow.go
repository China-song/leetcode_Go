package slidingwindow

/*
2760. Longest Even Odd Subarray With Threshold
*/
func longestAlternatingSubarray(nums []int, threshold int) (res int) {
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] > threshold || nums[i]%2 != 0 {
			i++
			continue
		}

		start := i
		i++
		for i < n && nums[i] <= threshold && nums[i]%2 != nums[i-1] {
			i++
		}
		// [start, i)
		res = max(res, i-start)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
3. Longest Substring Without Repeating Characters
*/
func lengthOfLongestSubstring(s string) (res int) {
	freq := [128]int{}
	left := 0
	for right, ch := range s {
		freq[ch]++
		for freq[ch] > 1 {
			freq[s[left]]--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

/*
209. Minimum Size Subarray Sum
*/
func minSubArrayLen(target int, nums []int) (res int) {
	n := len(nums)
	res = n + 1
	left := 0
	s := 0
	for right, num := range nums {
		s += num
		for s >= target {
			res = min(res, right-left+1)
			s -= nums[left]
			left++
		}
	}
	if res == n+1 {
		return 0
	} else {
		return res
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
713. Subarray Product Less Than K
*/
func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	if k <= 1 {
		return 0
	}
	prod := 1
	left := 0
	for right, x := range nums {
		prod *= x
		for prod >= k {
			prod /= nums[left]
			left++
		}
		// prod[left .. right] < k
		ans += right - left + 1
	}
	return ans
}

/*
689. 三个无重叠子数组的最大和
*/
func maxSumOfThreeSubarrays(nums []int, k int) (ans []int) {
	var sum1, maxSum1, maxSum1Idx int
	var sum2, maxSum12, maxSum12Idx1, maxSum12Idx2 int
	var sum3, maxSum123 int
	// [0, k-1], [k, 2*k-1], [2*k, 3*k-1]
	// 在已知前两个无重叠子数组的最大和maxSum12的情况下，遍历第三个sum3，以求三者最大和
	for i := 2 * k; i < len(nums); i++ {
		sum1 += nums[i-2*k]
		sum2 += nums[i-k]
		sum3 += nums[i]
		if i >= 3*k-1 {
			if sum1 > maxSum1 {
				maxSum1 = sum1
				maxSum1Idx = i - 3*k + 1
			}
			if maxSum1+sum2 > maxSum12 {
				maxSum12 = maxSum1 + sum2
				maxSum12Idx1, maxSum12Idx2 = maxSum1Idx, i-2*k+1
			}
			if maxSum12+sum3 > maxSum123 {
				maxSum123 = maxSum12 + sum3
				ans = []int{maxSum12Idx1, maxSum12Idx2, i - k + 1}
			}
			sum1 -= nums[i-3*k+1]
			sum2 -= nums[i-2*k+1]
			sum3 -= nums[i-k+1]
		}
	}
	return ans
}
