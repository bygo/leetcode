package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre int
var res int

func minDiffInBST(root *TreeNode) int {
	res = 1<<63 - 1
	pre = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		if pre != 0 {
			res = min(res, root.Val-pre)
		}
		pre = root.Val
		dfs(root.Right)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
