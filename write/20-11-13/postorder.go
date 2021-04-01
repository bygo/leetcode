package _0_11_12

func postorder(root *TreeNode) []int {
	var res []int
	postdfs(root,res)
	return res
}


func postdfs(root *TreeNode, res []int) {
	if root != nil {
		postdfs(root.Left, res)
		postdfs(root.Right,res)
		res = append(res, root.Val)
	}
}

func postorder_stack(root *TreeNode) []int {
	var res []int
	var stack = []*TreeNode{root}
	for 0 < len(stack) {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root.Left)
			root = root.Right
		}
		top := len(stack) - 1
		root = stack[top]
		stack = stack[:top]
	}

	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}

func postorder_morris_break(root *TreeNode) []int {
	var res []int
	for root != nil {
		res = append(res, root.Val)
		if root.Right == nil {
			root = root.Left
		} else {
			min := root.Right
			for min.Left != nil {
				min = min.Left
			}
			min.Left = root
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

func postorder_morris_keep(root *TreeNode) []int {
	var res []int
	for root != nil {
		if root.Right == nil {
			res = append(res, root.Val)
			root = root.Left
		} else {
			min := root.Right
			for min.Left != nil && min.Left != root {
				min = min.Left
			}
			if min.Left == nil {
				res = append(res, root.Val)
				min.Left = root
				root = root.Right
			} else {
				root = root.Left
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
