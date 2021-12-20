package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func distributeCoins(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := dfs(root.Left), dfs(root.Right)
	res += abs(l) + abs(r)
	return l + r + root.Val - 1
}

func abs(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}
