package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func findTilt(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) int {
	if root != nil {
		l := dfs(root.Left)
		r := dfs(root.Right)
		res += abs(l - r)
		return root.Val + l + r
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
