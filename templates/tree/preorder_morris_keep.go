package tree

func preorderMorrisKeep(root *TreeNode) *TreeNode {
	head := root
	var max *TreeNode

	for root != nil {
		if root.Left == nil {
			root = root.Right
		} else {
			//寻找左树最大节点
			max = root.Left
			for max.Right != nil && max != root {
				max = max.Right
			}

			if max.Right == nil {
				max.Right = root
				root = root.Left
			} else {
				root = root.Right
				max.Right = nil
			}
		}
	}
	return head
}
