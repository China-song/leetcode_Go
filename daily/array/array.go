package array

import (
	"math"
	"sort"
	"strconv"
)

/*
ThreeSumClosest16
16. 3Sum Closest
find three num in nums which sum close to target
*/
func ThreeSumClosest16(nums []int, target int) int {
	n := len(nums)
	// first, sort nums in increasing order
	sort.Ints(nums)

	closest := math.MaxInt32

	update := func(sum int) {
		if abs(sum-target) < abs(closest-target) {
			closest = sum
		}
	}

	// throughout first num
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a := nums[i]

		// throughout all b, c in nums[i+1, n-1]
		// find b, c from left and right: [i+1, n-1]
		for j, k := i+1, n-1; j < k; {
			b := nums[j]
			c := nums[k]

			sum := a + b + c
			if sum == target {
				return target
			}
			// try update
			update(sum)
			// scale range about b and c
			if sum < target {
				// b is not considered anymore because no third number(using this b) can meet their sum closer to target
				// move b to right to enlarge b so that enlarge sum
				for j+1 < k && nums[j+1] == nums[j] {
					j++
				}
				j++
			} else { // a + b + c > target
				// c is not considered anymore because no second number(using this c) can meet their sum closer to target
				// move c to left to reduce c so that reduce sum
				for k-1 > j && nums[k-1] == nums[k] {
					k--
				}
				k--
			}
		}
	}
	return closest

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
FindTheArrayConcVal2562
2562. Find the Array Concatenation Value
*/
func FindTheArrayConcVal2562(nums []int) int64 {
	var answer int64
	n := len(nums)
	for i, j := 0, n-1; i <= j; {
		if i == j {
			answer += int64(nums[i])
			break
		}
		concat, _ := strconv.Atoi(strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]))
		answer += int64(concat)
		i++
		j--
	}
	return answer
}

/*
307. Range Sum Query - Mutable
*/
type NumArray struct {
	nums []int
	tree []int // tree[i] is sum of nums[1..i]
}

func Constructor(nums []int) NumArray {
	a := NumArray{make([]int, len(nums)), make([]int, len(nums)+1)}
	for i, num := range nums {
		a.Update(i, num)
	}
	return a
}

func (a *NumArray) Update(index int, val int) {
	// 找到包含index的tree
	// tree是区间和 直接+delta
	delta := val - a.nums[index]
	a.nums[index] = val
	for i := index + 1; i < len(a.tree); i += i & -i {
		a.tree[i] += delta
	}
}

func (a *NumArray) prefixSum(i int) (sum int) {
	// 1..i 关键区求和 [1..x] [x+1..y] .. [z..i]
	/*
		8=[1..8]
		7=[1..4] [5..6] [7..7]
		6=[1..4] [5..6]
		5=[1..4] [5..5]
		4=[1..4]
		3=[1..2] [3..3]
		2=[1..2]
		1=[1..1]
	*/
	/*
	 */
	for ; i > 0; i -= i & -i {
		sum += a.tree[i]
	}
	return
}

func (a *NumArray) SumRange(left int, right int) (sum int) {
	// sum = presum[right+1] -
	sum = a.prefixSum(right+1) - a.prefixSum(left)
	return
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

/*
2760. Longest Even Odd Subarray With Threshold
*/
func LongestAlternatingSubarray(nums []int, threshold int) (res int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i]%2 != 0 {
			continue
		}
		cnt := 0
		x := 0
		for j := i; j < n; j++ {
			if nums[j] <= threshold && nums[j]%2 == x {
				cnt++
				x = x ^ 1
			} else {
				break
			}
		}
		res = max(res, cnt)
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
2342. Max Sum of a Pair With Equal Sum of Digits
*/
func maximumSum(nums []int) int {
	ans := -1
	mx := [82]int{}
	for _, num := range nums {
		s := 0
		for x := num; x > 0; x /= 10 {
			s += x % 10
		}
		if mx[s] > 0 {
			ans = max(ans, mx[s]+num)
		}
		mx[s] = max(mx[s], num)
	}
	return ans
}

/*
2661. First Completely Painted Row or Column
*/
func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	// rowCnt[i] == n 表示第i行涂满
	// colCnt[j] == m 表示第j列涂满
	rowCnt, colCnt := make([]int, m), make([]int, n)
	// pos[i]表示数字i的位置
	pos := make([]int, m*n)
	for i, row := range mat {
		for j, num := range row {
			pos[num] = i*m + j
		}
	}

	for k, num := range arr {
		i, j := pos[num]/m, pos[num]%n
		rowCnt[i]++
		colCnt[j]++
		if rowCnt[i] == n || colCnt[j] == m {
			return k
		}
	}
	return m*n - 1
}
