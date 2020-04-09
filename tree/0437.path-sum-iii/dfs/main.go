package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int
var s int

func pathSum(root *TreeNode, sum int) int {
	res = 0
	s = sum
	fall(root)
	return res
}

func fall(root *TreeNode) {
	if root == nil {
		return
	}
	fall(root.Left)
	fall(root.Right)
	dfs(root, 0)
}

func dfs(root *TreeNode, cur int) {
	if root != nil {
		cur += root.Val
		if s == cur {
			res++

		}
		dfs(root.Left, cur)
		dfs(root.Right, cur)
	}
}
