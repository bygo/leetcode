package tree

func preorderMorrisBreak(root *TreeNode) {
	var max *TreeNode

	for root != nil {
		if root.Left == nil {
			println(root.Val)
			root = root.Right
		} else {
			//寻找左树最大节点
			max = root.Left
			for max.Right != nil {
				max = max.Right
			}

			root.Right, max.Right = root.Left, root.Right
			root.Left = nil
		}
	}

}
