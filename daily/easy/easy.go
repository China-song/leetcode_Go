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

/*
CountPoints2103
2103. 环和杆
*/
func CountPoints2103(rings string) int {
	//res := 0
	//ring := make([][3]bool, 10)
	//color := map[uint8]int{'R': 0, 'G': 1, 'B': 2}
	//for i := 0; i < len(rings); i += 2 {
	//	ring[rings[i+1]-'0'][color[rings[i]]] = true
	//}
	//for i := 0; i < 10; i++ {
	//	if ring[i][0] && ring[i][1] && ring[i][2] {
	//		res++
	//	}
	//}
	//return res
	state := make([]int, 10)
	color := map[byte]int{'R': 1, 'G': 2, 'B': 4}
	n := len(rings)
	for i := 0; i < n; i += 2 {
		state[rings[i+1]-'0'] |= color[rings[i]]
	}
	res := 0
	for i := 0; i < 10; i++ {
		if state[i] == 7 {
			res++
		}
	}
	return res
}
