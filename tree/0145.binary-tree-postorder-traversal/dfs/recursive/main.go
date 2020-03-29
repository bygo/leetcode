package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func postorderTraversal(root *TreeNode) []int {
	res = []int{}
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		dfs(root.Right)
		res = append(res, root.Val)
	}
}
