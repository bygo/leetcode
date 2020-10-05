package tree

func inorderMorrisBreak(root *TreeNode) []int {
	var res []int
	var max *TreeNode

	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val) //中序遍历
			root = root.Right           //移动
		} else {
			max = root.Left //寻找左树最大值
			for max.Right != nil {
				max = max.Right
			}

			//root将在下一次输出
			max.Right = root                 //左树最大值连接root
			root, root.Left = root.Left, nil //移动左节点，砍左树
		}
	}
	return res
}
