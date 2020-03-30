package tree

func inorderMorrisKeep(root *TreeNode) []int {
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
				max.Right = root
				root = root.Left
			} else {
				res = append(res, root.Val) //中序输出
				root = root.Right
				max.Right = nil
			}
		}
	}

	return res
}
