package tree

func preorderMorrisBreak(root *TreeNode) []int {
	var max *TreeNode
	var res []int
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
				res = append(res, root.Val)
				max.Right = root.Right
				root = root.Left
			} else {
				root = root.Right
				max.Right = nil
			}
		}
	}
	return res
}
