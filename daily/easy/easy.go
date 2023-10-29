package easy

import (
	"sort"
	"strconv"
)

/*
SplitNum2578
2578. Split With Minimum Sum
*/
func SplitNum2578(num int) int {
	// sort digits of num
	stnum := []byte(strconv.Itoa(num))
	sort.Slice(stnum, func(i, j int) bool {
		return stnum[i] < stnum[j]
	})
	// sorted: {1, 2 ,4 ,5}

	// distribute to 2 num   [1,4] [2,5]
	// transfer to num 14 25
	num1, num2 := 0, 0
	for i := 0; i < len(stnum); i++ {
		if i%2 == 0 {
			num1 = num1*10 + int(stnum[i]-'0')
		} else {
			num2 = num2*10 + int(stnum[i]-'0')
		}
	}

	return num1 + num2
}

func categorizeBox(length int, width int, height int, mass int) string {
	bulky := 0
	heavy := 0
	var volume int64
	volume = int64(length) * int64(width) * int64(height)
	if length >= 1e4 || width >= 1e4 || height >= 1e4 || volume >= int64(1e9) {
		bulky = 1
	}
	if mass >= 100 {
		heavy = 1
	}

	ans := [2][2]string{{"Neither", "Heavy"}, {"Bulky", "Both"}}
	return ans[bulky][heavy]
}
