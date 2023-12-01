package string

import (
	"sort"
	"strconv"
	"strings"
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

/*
1410. HTML Entity Parser
*/
func entityParser(s string) (ans string) {
	return strings.NewReplacer(`&quot;`, `"`, `&apos;`, `'`, `&gt;`, `>`, `&lt;`, `<`, `&frasl;`, `/`, `&amp;`, `&`).Replace(s)
}

/*
828. Count Unique Characters of All Substrings of a Given String
*/
func uniqueLetterString(s string) (ans int) {
	//
	idx := map[rune][]int{}
	for i, c := range s {
		idx[c] = append(idx[c], i)
	}
	for _, arr := range idx {
		arr = append(append([]int{-1}, arr...), len(s))
		for i := 1; i < len(arr)-1; i++ {
			// j i k
			// j+1 j+2 .. j+n=i
			ans += (arr[i] - arr[i-1]) * (arr[i+1] - arr[i])
		}
	}
	return ans
}

/*
1657. Determine if Two Strings Are Close
*/
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	// 记录word中出现过的字母的出现频率
	freq1 := make([]int, 26)
	freq2 := make([]int, 26)
	s, t := 0, 0

	for i := 0; i < len(word1); i++ {
		freq1[word1[i]-'a']++
		s = s | (1 << (word1[i] - 'a'))

		freq2[word2[i]-'a']++
		t = t | (1 << (word2[i] - 'a'))
	}
	if s^t != 0 {
		return false
	}
	sort.Ints(freq1)
	sort.Ints(freq2)
	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}
	return true
}
