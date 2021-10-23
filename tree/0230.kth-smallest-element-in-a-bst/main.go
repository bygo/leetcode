package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var res int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			k--
			if k == 0 {
				res = root.Val
				return
			}
			dfs(root.Right)
		}
	}
	dfs(root)
	return res
}
