package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int
var pre *TreeNode

func getMinimumDifference(root *TreeNode) int {
	res = 1<<63 - 1
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, f int) {
	if root != nil {
		dfs(root.Left, root.Val)
		if pre != nil {
			r := abs(root.Val - pre.Val)
			if r < res {
				res = r
			}
		}
		pre = root
		dfs(root.Right, root.Val)
	}
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
