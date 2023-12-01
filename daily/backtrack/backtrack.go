package backtrack

/*
17. Letter Combinations of a Phone Number
*/
func letterCombinations(digits string) (res []string) {
	mp := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var backtrack func(int, int, string)
	backtrack = func(i int, n int, s string) {
		if i == n {
			res = append(res, s)
		}
		for _, c := range mp[digits[i]] {
			backtrack(i+1, n, s+string(c))
		}
	}
	backtrack(0, len(digits), "")
	return res
}

/*
1457. Pseudo-Palindromic Paths in a Binary Tree
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) (ans int) {
	var a [10]int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			cnt := 0
			for i := 1; i <= 9; i++ {
				if a[i]%2 == 1 {
					cnt++
				}
			}
			if cnt <= 1 {
				ans++
			}
			return
		}

		if node.Left != nil {
			a[node.Left.Val]++
			dfs(node.Left)
			a[node.Left.Val]--
		}
		if node.Right != nil {
			a[node.Right.Val]++
			dfs(node.Right)
			a[node.Right.Val]--
		}
	}
	a[root.Val]++
	dfs(root)
	return ans
}
