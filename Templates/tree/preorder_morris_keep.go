package tree

func preorderMorrisKeep(root *TreeNode) []int {
	var res []int
	var max *TreeNode
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val) //左节点为空 前序遍历
			root = root.Right           //移动到右节点
		} else {
			max = root.Left //寻找左树最大值
			for max.Right != nil && max.Right != root.Right {
				max = max.Right
			}

			//前序指针处理
			if max.Right == nil {
				res = append(res, root.Val) //前序遍历
				max.Right = root.Right      //指向
				root = root.Left            //移动到左节点
			} else { //已指向
				root = root.Right //跳跃
				max.Right = nil   //删指向
			}
		}
	}
	return res
}
