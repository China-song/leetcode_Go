package array

import "testing"

func TestThreeSumClosest16(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		ans    int
	}{
		{
			nums:   []int{1, 1, 1, 0},
			target: -100,
			ans:    2,
		},
	}

	for _, cas := range cases {
		res := ThreeSumClosest16(cas.nums, cas.target)
		if res != cas.ans {
			t.Fatalf("expected %d, but get %d", cas.ans, res)
		}
	}

}

func TestLongestAlternatingSubarray(t *testing.T) {
	LongestAlternatingSubarray([]int{2, 3, 4, 5}, 4)
}
