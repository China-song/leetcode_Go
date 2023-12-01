package orderedset

import "github.com/emirpasic/gods/trees/redblacktree"

type RangeModule struct {
	*redblacktree.Tree
}

func Constructor() RangeModule {
	return RangeModule{redblacktree.NewWithIntComparator()}
}

func (t *RangeModule) AddRange(left int, right int) {
	if node, ok := t.Floor(left); ok {
		// li <= left < l i+1
		r := node.Value.(int)
		if right <= r {
			// li <= left < right <= ri	不需要添加
			return
		}
		// ri < right
		if left <= r {
			// li <= left <= ri < right
			// 有交集
			// 去掉[li, ri) 添加[li, right)
			t.Remove(node.Key.(int))
			left = node.Key.(int)
		}
	}
	// left最小 或者 空区间
	for node, ok := t.Ceiling(left); ok && node.Key.(int) <= right; node, ok = t.Ceiling(left) {
		right = max(right, node.Value.(int))
		t.Remove(node.Key)
	}

	// 添加[left, right)
	t.Put(left, right)
}

func (t *RangeModule) QueryRange(left int, right int) bool {
	node, ok := t.Floor(left)
	return ok && node.Value.(int) >= right
}

func (t *RangeModule) RemoveRange(left int, right int) {
	if node, ok := t.Floor(left); ok {
		l, r := node.Key.(int), node.Value.(int)
		if right <= r {
			if left == l {
				t.Remove(l)
			} else {
				// l < left < right <= r
				node.Value = left
			}
			if right != r {
				t.Put(right, r)
			}
			return
		}
		if left < r {
			if left == l {
				t.Remove(l)
			} else {
				node.Value = left
			}
		}
	}

	for node, ok := t.Ceiling(left); ok && node.Key.(int) < right; node, ok = t.Ceiling(left) {
		r := node.Value.(int)
		t.Remove(node.Key)
		if r > right {
			t.Put(right, r)
			break
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */

/*
56. Merge Intervals
*/
func merge(intervals [][]int) (res [][]int) {
	t := redblacktree.NewWithIntComparator()
	for _, interval := range intervals {
		left, right := interval[0], interval[1]
		// deal left first
		if node, ok := t.Floor(left); ok {
			// l1 <= left < l2
			l := node.Key.(int)
			r := node.Value.(int)
			if right <= r {
				// l <= left <= right <= r	no need to merge this interval
				continue
			}
			// l <= r < right
			if left <= r {
				// l <= left <= r < right	merge
				t.Remove(l)
				left = l
			}
		}
		// finish left or left is the smallest
		// deal right
		// left <= l <= r
		for node, ok := t.Ceiling(left); ok && right >= node.Key.(int); node, ok = t.Ceiling(left) {
			r := node.Value.(int)
			t.Remove(node.Key)
			right = max(right, r)
		}

		// finally, put it
		t.Put(left, right)
	}
	for _, left := range t.Keys() {
		right, _ := t.Get(left)
		res = append(res, []int{left.(int), right.(int)})
	}
	return
}
