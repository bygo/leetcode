package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var last *TreeNode

	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			last = root.Left
			for last.Right != nil {
				last = last.Right
			}
			last.Right, root, root.Left = root, root.Left, nil
		}
	}
	return res
}
