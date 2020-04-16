package tree

func inorderMorrisBreak(root *TreeNode) []int {
	var res []int
	var max *TreeNode
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val) //中序遍历
			root = root.Right           //链表移动
		} else {
			max = root.Left //寻找左树最大值
			for max.Right != nil {
				max = max.Right
			}

			//中序指针处理，root将在下一次循环输出
			max.Right = root                 //左树最大值连接 root
			root, root.Left = root.Left, nil //移到左节点，砍 root 左树
		}
	}
	return res
}
