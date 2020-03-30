package tree

func preorderMorrisBreak(root *TreeNode) []int {
	var res []int
	var max *TreeNode
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val) //前序输出
			root = root.Right           //链表移动
		} else {
			max = root.Left //寻找左树最大节点
			for max.Right != nil {
				max = max.Right
			}

			//前序指针处理，root将在下一次循环输出
			root.Right, max.Right = root.Left, root.Right //转为以Right为方向的链表
			root.Left = nil                               //砍左树
		}
	}
	return res
}
