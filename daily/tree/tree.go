package tree

/*
104. Maximum Depth of Binary Tree
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		ans++
		var tmp []*TreeNode
		for _, node := range q {
			// 根据上一层结点得到下一层结点
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		q = tmp
	}
	return ans
}
