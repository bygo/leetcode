package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int
var l, r int

func rangeSumBST(root *TreeNode, L int, R int) int {
	res = 0
	l, r = L, R
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		if l <= root.Val && root.Val <= r {
			res += root.Val
		}
		if l < root.Val {
			dfs(root.Left)
		}
		if root.Val < r {
			dfs(root.Right)
		}
	}
}
