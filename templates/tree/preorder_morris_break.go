package tree

func preorderMorrisBreak(root *TreeNode) []int {
	var res []int
	var max *TreeNode
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right //其实只是单向链表了
		} else {
			//寻找左树最大节点
			max = root.Left
			for max.Right != nil {
				max = max.Right
			}

			root.Right, max.Right = root.Left, root.Right //即转为以Right为方向的链表
			root.Left = nil                               //砍左树
		}
	}
	return res
}
