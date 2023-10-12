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
