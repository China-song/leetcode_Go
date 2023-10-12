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
