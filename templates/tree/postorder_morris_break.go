package tree

func postorderMorrisBreak(root *TreeNode) []int {
	var res []int
	var min *TreeNode
	for root != nil {
		res = append(res, root.Val) //后序遍历
		if root.Right == nil {
			root = root.Left //移动
		} else {
			min = root.Right //寻找右树最小节点
			for min.Left != nil && min.Left != root {
				min = min.Left
			}

			//后序指针处理
			min.Left = root.Left
			root, root.Right = root.Right, nil
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
