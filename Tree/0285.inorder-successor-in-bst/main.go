package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var max *TreeNode
	for root != nil {
		if root.Left == nil {
			if root == p {
				root = root.Right
				if root == nil {
					return nil
				}
				for root.Left != nil {
					root = root.Left
				}
				return root
			}
			root = root.Right
		} else {
			max = root.Left
			for max.Right != nil {
				max = max.Right
			}

			max.Right = root
			root, root.Left = root.Left, nil

		}
	}
	return nil
}
