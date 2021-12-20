package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var last *TreeNode
var res int

func minDiffInBST(root *TreeNode) int {
	last = nil
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		if last != nil {
			r := root.Val - last.Val
			if res == 0 || r < res {
				res = r
			}
		}
		last = root
		dfs(root.Right)
	}
}
