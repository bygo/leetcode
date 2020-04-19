package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	dfs(root)
	return root
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Right)
		sum += root.Val
		root.Val = sum
		dfs(root.Left)
	}
}
