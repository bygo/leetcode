package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int
var last *TreeNode

func getMinimumDifference(root *TreeNode) int {
	res = 0
	last = nil
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
