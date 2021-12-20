package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var max *TreeNode

	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			max = root.Left
			for max.Right != nil && max.Right != root {
				max = max.Right
			}

			if max.Right == nil {
				// 未连接,连接到root
				max.Right = root
				root = root.Left
			} else {
				// 已连接,断开连接
				max.Right = nil
				res = append(res, root.Val)
				root = root.Right
			}
		}
	}
	return res
}
