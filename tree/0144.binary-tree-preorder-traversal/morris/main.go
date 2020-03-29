package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var max *TreeNode
	var res []int
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			max = root.Left
			for max.Right != nil {
				max = max.Right
			}

			root.Right, max.Right = root.Left, root.Right
			root.Left = nil
		}
	}
	return res
}
