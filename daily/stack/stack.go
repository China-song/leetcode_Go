package stack

import "math"

// NextGreaterElement496
// 496. 下一个更大元素 I
func NextGreaterElement496(nums1 []int, nums2 []int) []int {
	// 先遍历nums2 找到每个元素key的下一个更大元素value 存到map中

	// map: 存放nums2每个元素的下一个更大元素
	nextGreat := make(map[int]int)
	// 单调栈
	stack := make([]int, 0)
	// 逆序存放
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] < nums2[i] {
			// pop
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			nextGreat[nums2[i]] = -1
		} else {
			nextGreat[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	// 再遍历nums1 map.get
	ans := make([]int, len(nums1))
	for i, num := range nums1 {
		ans[i] = nextGreat[num]
	}

	return ans
}

/*
StockSpanner
901. 股票价格跨度
*/
type StockSpanner struct {
	// 单调栈
	// stack[i][0]存放price stack[i][1]存放索引
	stack [][2]int
	idx   int // 表示当前price是第几天的price
}

func Constructor() StockSpanner {
	// 第一天price前面没有比它低的 所以栈初始放个最大值 索引为-1
	return StockSpanner{stack: [][2]int{{math.MaxInt, -1}}, idx: 0}
}

func (this *StockSpanner) Next(price int) int {
	var span int
	// 维持单调栈
	for this.stack[len(this.stack)-1][0] <= price {
		// pop
		this.stack = this.stack[:len(this.stack)-1]
	}
	// 栈顶价格比price大: 找到price之前的更大价格
	span = this.idx - this.stack[len(this.stack)-1][1]

	// 存入当前price以及它的索引
	this.stack = append(this.stack, [2]int{price, this.idx})
	// 为后续price更新索引
	this.idx++
	return span
}
