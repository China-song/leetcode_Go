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

/*
FindMaximumXOR421
421. Maximum XOR of Two Numbers in an Array
*/
const highBit = 30

type trie struct {
	left, right *trie
}

func (t *trie) add(num int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.left == nil {
				cur.left = &trie{}
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &trie{}
			}
			cur = cur.right
		}
	}
}

func (t *trie) check(num int) (x int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1 // ai的当前位
		if bit == 0 {
			if cur.right != nil {
				x = x<<1 + 1 // x当前位可取1
				cur = cur.right
			} else {
				x = x << 1
				cur = cur.left
			}
		} else {
			if cur.left != nil {
				x = x<<1 + 1 // x当前位可取1
				cur = cur.left
			} else {
				x = x << 1
				cur = cur.right
			}
		}
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindMaximumXOR421(nums []int) (x int) {
	/*
		方法二: 字典树
	*/
	// 构造字典树
	root := &trie{}
	for i := 1; i < len(nums); i++ {
		root.add(nums[i-1])
		x = max(x, root.check(nums[i]))
	}
	return
	/*
		方法一：哈希表

		const highBit = 30 // 最高位的二进制位编号为 30
		for k := highBit; k >= 0; k-- {
			seen := map[int]bool{}
			for _, num := range nums {
				seen[num>>k] = true
			}
			xNext := (x << 1) + 1
			found := false

			for _, num := range nums {
				if seen[num>>k^xNext] {
					found = true
					break
				}
			}

			if found {
				x = xNext
			} else {
				x = xNext - 1
			}
		}
		return

	*/
}
