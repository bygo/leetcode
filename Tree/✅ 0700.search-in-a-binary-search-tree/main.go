package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val < val {
			root = root.Right
		} else if val < root.Val {
			root = root.Left
		} else {
			return root
		}
	}
	return nil
}
