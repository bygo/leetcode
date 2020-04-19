package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res *TreeNode
var target int

func searchBST(root *TreeNode, val int) *TreeNode {
	res = nil
	target = val
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil && res == nil {
		if root.Val == target {
			res = root
		}
		dfs(root.Left)
		dfs(root.Right)
	}
}
