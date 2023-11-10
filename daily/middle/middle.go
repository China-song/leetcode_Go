package middle

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"sort"
	"strings"
)

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

/*
TopStudents2512
2512. Reward Top K Students
*/
func TopStudents2512(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	// first, make map[feedback]point
	positive_map := make(map[string]int)
	for _, positive := range positive_feedback {
		positive_map[positive] = 3
	}
	negative_map := make(map[string]int)
	for _, negative := range negative_feedback {
		negative_map[negative] = -1
	}

	// construct a points array about students
	tree := redblacktree.NewWith(func(a, b interface{}) int {
		s1 := a.(Student)
		s2 := b.(Student)
		if s1.Point >= s2.Point {
			if s1.Point > s2.Point {
				return -1
			} else if s1.Id < s2.Id {
				return -1
			} else if s1.Id == s2.Id {
				return 0
			} else {
				return 1
			}
		} else {
			return 1
		}
	})
	n := len(report)
	for i := 0; i < n; i++ {
		points := 0
		feedbacks := strings.Split(report[i], " ")
		for _, feedback := range feedbacks {
			if positive_map[feedback] != 0 {
				points += 3
			} else if negative_map[feedback] != 0 {
				points -= 1
			}
		}
		tree.Put(Student{Point: points, Id: student_id[i]}, nil)
	}

	topK := make([]int, 0)
	students := tree.Keys()
	for i, student := range students {
		id := student.(Student).Id
		topK = append(topK, id)
		if i == k-1 {
			break
		}
	}
	return topK
}

type Student struct {
	Point int
	Id    int
}

/*
MaxArea1465
1465. Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts
*/
func MaxArea1465(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	// sort in increasing order
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	maxH, maxW := 0, 0
	for i := 0; i < len(horizontalCuts); i++ {
		if i == 0 {
			maxH = horizontalCuts[0]
		} else {
			maxH = max(maxH, horizontalCuts[i]-horizontalCuts[i-1])
		}
	}
	maxH = max(maxH, h-horizontalCuts[len(horizontalCuts)-1])
	for i := 0; i < len(verticalCuts); i++ {
		if i == 0 {
			maxW = verticalCuts[0]
		} else {
			maxW = max(maxW, verticalCuts[i]-verticalCuts[i-1])
		}
	}
	maxW = max(maxW, w-verticalCuts[len(verticalCuts)-1])

	return maxH * maxW
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func tupleSameProduct(nums []int) int {
	n := len(nums)
	if n < 4 {
		return 0
	}
	sort.Ints(nums)
	ans := 0
	a, c, d, b := 0, 1, 2, 3
	for ; a < n; a++ {
		for ; c < n; c++ {
			for ; d < n; d++ {
				for ; b < n; b++ {
					left := nums[a] * nums[b]
					right := nums[c] * nums[d]
					if left < right {
						continue
					} else if left == right {
						ans++
					} else {
						break
					}
				}
			}
		}
	}
	return ans * 8
}

/*
HIndex274
274. H-Index
*/
func HIndex274(citations []int) int {
	// 方法一
	//sort.Ints(citations)
	//h := 0 // 当前h-index
	//for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
	//	h++
	//}
	//return h

	// 方法三：二分搜索
	// [left, right]
	left, right := 0, len(citations)
	var mid int
	for left < right {
		mid = (left + right + 1) >> 1
		cnt := 0
		for _, v := range citations {
			if v >= mid {
				cnt++
			}
		}
		if cnt >= mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

/**
* Definition for a Node.
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for i, node := range tmp {
			if (i + 1) < len(tmp) {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}

/*
FindRepeatedDnaSequences187
187. Repeated DNA Sequences
*/
func FindRepeatedDnaSequences187(s string) (ans []string) {
	// 哈希表 + 滑动窗口 + 位运算
	const L = 10
	mp := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

	n := len(s)
	if n <= L {
		return
	}
	x := 0
	for _, ch := range s[:L-1] {
		x = x<<2 | mp[byte(ch)]
	}

	cnt := map[int]int{}
	for i := 0; i <= n-L; i++ {
		x = (x<<2 | mp[s[i+L-1]]) & ((1 << 20) - 1)
		cnt[x]++
		if cnt[x] == 2 {
			ans = append(ans, s[i:i+L])
		}
	}
	return
}
