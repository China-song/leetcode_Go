package greedy

import "sort"

/*
AvoidFlood1488
1488. Avoid Flood in The City
*/
func AvoidFlood1488(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	lakes := make(map[int]int)
	dryDays := make([]int, 0)
	for i := 0; i < n; i++ {
		ans[i] = 1
	}

	for i, rain := range rains {
		if rain == 0 {
			// store dry day
			dryDays = append(dryDays, i)
		} else {
			// rain > 0 indicates a lake
			// there will be rain over the lake
			ans[i] = -1
			if day, ok := lakes[rain]; ok {
				// lake is full of water
				// choose a day which is close and after rain
				idx := sort.SearchInts(dryDays, day)
				if idx == len(dryDays) {
					return []int{} // no day to dry
				}
				ans[dryDays[idx]] = rain // indicates the day that dry lake

				// remove this day from dryDays
				copy(dryDays[idx:], dryDays[idx+1:])
				dryDays = dryDays[:len(dryDays)-1]
			}
			lakes[rain] = i
		}
	}

	return ans
}
