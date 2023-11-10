package string

import (
	"strconv"
)

/*
CountSeniors2678
2678. Number of Senior Citizens
*/
func CountSeniors2678(details []string) int {
	ans := 0
	for _, detail := range details {
		age, _ := strconv.Atoi(detail[11:13])
		if age > 60 {
			ans++
		}
	}
	return ans
}

/*
MaxProduct318
318. Maximum Product of Word Lengths
*/
func MaxProduct318(words []string) int {
	// meet 与 met 具有相同的mask，只需记录长度更大的mask与它对应的长度
	masks := map[int]int{}
	for _, word := range words {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		if len(word) > masks[mask] {
			masks[mask] = len(word)
		}
	}
	res := 0
	for maskX, lenX := range masks {
		for maskY, lenY := range masks {
			if maskX&maskY == 0 && lenX*lenY > res {
				res = lenX * lenY
			}
		}
	}
	return res
	//masks := make([]int, len(words))
	//for i, word := range words {
	//	for _, ch := range word {
	//		masks[i] |= 1 << (ch - 'a')
	//	}
	//}
	//res := 0
	//for i, x := range masks {
	//	for j, y := range masks[:i] {
	//		if x&y == 0 && len(words[i])*len(words[j]) > res {
	//			res = len(words[i]) * len(words[j])
	//		}
	//	}
	//}
	//return res
}
