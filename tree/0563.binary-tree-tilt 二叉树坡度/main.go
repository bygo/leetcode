package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	var res int
	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if n != nil {
			l := dfs(n.Left)
			r := dfs(n.Right)
			res += abs(l - r)
			return n.Val + l + r
		}
		return 0
	}
	dfs(root)
	return res
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
