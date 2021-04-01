package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	var res int
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, cur int, res *int) {
	if root != nil {
		cur <<= 1
		cur += root.Val
		if root.Left == nil && root.Right == nil {
			*res += cur
		}
		dfs(root.Left, cur, res)
		dfs(root.Right, cur, res)
	}
}
