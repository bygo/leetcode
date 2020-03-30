package tree

func preorderMorrisKeep(root *TreeNode) []int {
	var res []int
	var max *TreeNode
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val) //左节点为空 直接输出 (前序遍历),
			root = root.Right           //移动到右节点
		} else {
			max = root.Left //找左树最大值
			for max.Right != nil && max.Right != root {
				max = max.Right
			}

			if max.Right == nil { //最大值没指向root.Right
				res = append(res, root.Val) //指向了root.Right时 直接输出，(前序遍历)
				max.Right = root.Right
				root = root.Left //移动到左节点
			} else { //已指向
				root = root.Right //跳跃
				max.Right = nil   //删指向
			}
		}
	}
	return res
}
