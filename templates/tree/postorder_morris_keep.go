package tree

func postorderMorrisKeep(root *TreeNode) []int {
	var res []int
	var min *TreeNode
	for root != nil {
		if root.Right == nil {
			res = append(res, root.Val) //后序遍历
			root = root.Left            //移动
		} else {
			min = root.Right //寻找右树最小节点
			for min.Left != nil && min.Left != root.Left {
				min = min.Left
			}

			if min.Left == nil {
				res = append(res, root.Val) //后序遍历
				min.Left = root.Left
				root = root.Right
			} else {
				root = root.Left //跳跃
				min.Left = nil
			}
		}
	}

	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
