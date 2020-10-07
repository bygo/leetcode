package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	var res int
	dfs(root, root.Val, &res)
	return res
}

func dfs(root *TreeNode, parent int, res *int) bool {
	if root == nil {
		return true
	}

	l := dfs(root.Left, root.Val, res)
	r := dfs(root.Right, root.Val, res)

	if l || r {
		return false
	}
	*res ++
	return root.Val == parent
}
