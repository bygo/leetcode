package tree

func inorderMorrisBreak(root *TreeNode) []int {
	var res []int
	var max *TreeNode
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
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
	return res
}
