package binarysearch

/*
HIndex275
275. H-Index II
*/
func HIndex275(citations []int) int {
	n := len(citations)
	left, right := 0, len(citations)
	var mid int
	for left < right {
		mid = (left + right + 1) >> 1
		if citations[n-mid] >= mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
