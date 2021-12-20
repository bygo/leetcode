package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	var res int
	dfs(root, root.Val, &res)
	return res
}

func dfs(root *TreeNode, max int, res *int) {
	if root != nil {
		if max <= root.Val {
			max = root.Val
			*res++
		}
		dfs(root.Left, max, res)
		dfs(root.Right, max, res)
	}
}
