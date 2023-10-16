package bit

/*
SingleNumber137
137. Single Number II
*/
func SingleNumber137(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		// accelerate ith bit of nums
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		ans |= (total % 3) << i
	}
	return int(ans)
}

/*
SingleNumber260
260. Single Number III
*/
func SingleNumber260(nums []int) []int {
	x1 := int32(0)
	x2 := int32(0)
	x := int32(0)
	for _, num := range nums {
		x ^= int32(num)
	}
	// x = x1 ^ x2
	test := x & (-x)
	// test: 0...010...0	indicate lowest digit 1 position of x
	// so x1:.....1.....
	//    x2 .....0.....

	for _, num := range nums {
		if (int32(num) & test) != int32(0) {
			x1 ^= int32(num)
		} else {
			x2 ^= int32(num)
		}
	}

	return []int{int(x1), int(x2)}
}
