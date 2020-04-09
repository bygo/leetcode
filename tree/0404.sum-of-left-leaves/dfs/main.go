package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func sumOfLeftLeaves(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res += root.Left.Val
		return
	}
	dfs(root.Left)
	dfs(root.Right)
}
