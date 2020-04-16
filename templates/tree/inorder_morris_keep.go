package tree

func inorderMorrisKeep(root *TreeNode) []int {
	var res []int
	var max *TreeNode
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val) //中序遍历
			root = root.Right
		} else {
			max = root.Left //寻找左树最大值
			for max.Right != nil && max.Right != root {
				max = max.Right
			}
			//中序指针处理
			if max.Right == nil {
				max.Right = root //左树最大值指向 root
				root = root.Left //移动到左节点
			} else {
				res = append(res, root.Val) //中序遍历
				root = root.Right           //跳跃
				max.Right = nil             //删指向
			}
		}
	}
	return res
}
