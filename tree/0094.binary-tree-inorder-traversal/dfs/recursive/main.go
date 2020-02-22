package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
var res []int

func inorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
}
