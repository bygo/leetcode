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
	if root != nil {
		if root.Left == nil && root.Right == nil {
			if depth < min {
				min = depth
			}
			return
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
}
