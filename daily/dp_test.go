package daily

import "testing"

func TestMaxProfit123(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	expectedAnswer := 0
	maxProfit := MaxProfit123(prices)
	if maxProfit != expectedAnswer {
		t.Fatalf("answer is %d, but get %d", expectedAnswer, maxProfit)
	}
}

func TestMaxProfit309(t *testing.T) {
	cases := []struct {
		Input  []int
		Output int
	}{
		{
			Input:  []int{1, 2, 3, 0, 2},
			Output: 3,
		},
		{
			Input:  []int{1},
			Output: 0,
		},
	}

	for _, cas := range cases {
		output := MaxProfit309(cas.Input)
		if output != cas.Output {
			t.Fatalf("answer is %d, but get %d", cas.Output, output)
		}
	}
}
