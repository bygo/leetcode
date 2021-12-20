package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func longestConsecutive(root *TreeNode) int {
	res = 0
	dfs(root, 0, 1)
	return res
}

func dfs(root *TreeNode, parent, count int) {
	if root != nil {
		if parent+1 != root.Val {
			count = 1
		}
		dfs(root.Left, root.Val, count+1)
		dfs(root.Right, root.Val, count+1)

		if res < count {
			res = count
		}
	}
}
