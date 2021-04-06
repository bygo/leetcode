package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var min int

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min = 1<<63 - 1
	dfs(root, 1)
	return min
}

func dfs(root *TreeNode, depth int) {
	if root.Left == nil && root.Right == nil {
		if depth < min {
			min = depth
		}
		return
	}
	if root.Left != nil {
		dfs(root.Left, depth+1)
	}
	if root.Right != nil {
		dfs(root.Right, depth+1)
	}
}
