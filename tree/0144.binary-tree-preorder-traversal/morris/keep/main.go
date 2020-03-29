package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) {
	var max *TreeNode
	for root != nil {
		if root.Left == nil {
			println(root.Val)
			root = root.Right
		} else {
			max = root.Left
			for max.Right != nil && max.Right != root {
				max = max.Right
			}

			if max.Right == nil {
				println(root.Val)
				max.Right = root.Right
				root = root.Left
			} else {
				root = root.Right
				max.Right = nil
			}
		}
	}
}
