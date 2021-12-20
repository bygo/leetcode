package tree

func preorderMorrisBreak(root *TreeNode) []int {
	var res []int
	var max *TreeNode
	for root != nil {
		res = append(res, root.Val) //前序遍历
		if root.Left == nil {
			root = root.Right //移动
		} else {
			max = root.Left //找左树最大节点
			for max.Right != nil {
				max = max.Right
			}

			//前序指针处理
			max.Right = root.Right
			root, root.Left = root.Left, nil //转为以Right为方向的单向树
		}
	}
	return res
}
