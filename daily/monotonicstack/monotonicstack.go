package monotonicstack

/*
739. Daily Temperatures
*/
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	var stk []int // 存下标

	// 方法一 倒序
	/*
		// 倒序 因为只看下一更大 倒序并存储能加快速度 使用栈因为先看后一元素 栈后进先出
		// 栈中存放候选元素（可能的更大温度），且为了增加搜索速度，设置为单调栈
		for i := n - 1; i >= 0; i-- {
			t := temperatures[i]
			for len(stk) > 0 && temperatures[stk[len(stk)-1]] <= t {
				stk = stk[:len(stk)-1]
			}
			if len(stk) > 0 {
				ans[i] = stk[len(stk)-1] - i
			}
			stk = append(stk, i)
		}
	*/

	// 方法二 正序
	// 每个温度都有可能作为 之前温度 的 "下一更大温度"
	// 栈中存放 未找到下一更大 的 之前温度
	for i, t := range temperatures {
		// 先给之前温度 找 下一更大
		for len(stk) > 0 && t > temperatures[stk[len(stk)-1]] {
			j := stk[len(stk)-1]
			ans[j] = i - j
			stk = stk[:len(stk)-1]
			// 这一操作保证了栈的单调性 栈顶是最小的 如果当前温度不能作为栈顶的下一更大，自然也不能作为其它的下一更大
		}
		// 最后 当前温度 也作为 之前温度 存放
		stk = append(stk, i)
	}
	return ans
}

/*
1019. Next Greater Node In Linked List
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) (ans []int) {
	for i, node := 0, head; node != nil; i, node = i+1, node.Next {
		val := node.Val
		cur := node.Next
		for cur != nil && cur.Val <= val {
			cur = cur.Next
		}
		if cur == nil {
			ans = append(ans, 0)
		} else {
			ans = append(ans, cur.Val)
		}
	}
	return ans
}

/*
907. Sum of Subarray Minimums
*/
func sumSubarrayMins(arr []int) (ans int) {
	// 对每个元素arr[i] 找以它为最小元素的子数组个数
	// 思考：比arr[i]小的元素不能包括，
	// 那相等的呢 如果包括相等的元素，那么会存在一个包括多个相同元素的子数组
	// 对每个元素计算子数组个数时，会将那个大子数组计算多次，实际只需要一次
	// 在左侧找小于它的最近元素 在右侧找小于等于它的最近元素
	n := len(arr)
	left := make([]int, n)
	st := []int{-1} // 单调栈 存放元素索引
	// 遍历到arr[i]时 找左侧最近小 的数 以 前一个数为基准
	// 比前一个数>=的数是不需要看的
	for i, x := range arr {
		for len(st) > 1 && arr[st[len(st)-1]] >= x {
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	// 从最右边开始找
	right := make([]int, n)
	st = []int{n}
	for i := n - 1; i >= 0; i-- {
		for len(st) > 1 && arr[st[len(st)-1]] > arr[i] {
			st = st[:len(st)-1]
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}

	for i, x := range arr {
		ans += x * ((i - left[i]) * (right[i] - i))
	}
	return ans % (1e9 + 7)
}
