package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, -1<<63, 1<<63-1)
}

func dfs(root *TreeNode, min, max int) bool {
	return root == nil || min < root.Val && root.Val < max &&
		dfs(root.Left, min, root.Val) &&
		dfs(root.Right, root.Val, max)
}
