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
