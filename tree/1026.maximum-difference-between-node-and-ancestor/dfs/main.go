package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	var res int
	dfs(root, root.Val, root.Val, &res)
	return res
}

func dfs(root *TreeNode, min, max int, res *int) {
	if root != nil {
		if root.Val < min {
			min = root.Val
		}
		if max < root.Val {
			max = root.Val
		}
		if *res < max-min {
			*res = max - min
		}

		dfs(root.Left, min, max, res)
		dfs(root.Right, min, max, res)
	}
}
