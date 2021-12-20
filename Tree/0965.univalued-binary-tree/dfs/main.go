package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var cur int
var res bool

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res = true
	cur = root.Val
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		if root.Val != cur {
			res = false
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}
}
