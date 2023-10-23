package string

import "strconv"

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
