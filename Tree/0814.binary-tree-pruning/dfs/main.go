package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left = pruneTree(root.Left)
		root.Right = pruneTree(root.Right)
		if root.Left == nil && root.Right == nil && root.Val == 0 {
			return nil
		}
	}
	return root
}
