package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil { //面向测试用例
		return true
	}
	return bfs(root, -1<<63, 1<<63-1)
}

func bfs(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || max <= root.Val {
		return false
	}
	return bfs(root.Left, min, root.Val) && bfs(root.Right, root.Val, max)
}
