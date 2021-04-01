package _0_11_12

func preoder(root *TreeNode) []int {
	var res []int
	pdfs(root, &res)
	return res
}

func pdfs(root *TreeNode, res *[]int) {
	if root != nil {
		pdfs(root.Left, res)
		*res = append(*res, root.Val)
		pdfs(root.Right, res)
	}
}

func preorder_stack(root *TreeNode) []int {
	var res []int
	var stack = []*TreeNode{root}
	for 0 < len(stack) {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root.Right)
			root = root.Left
		}

		top := len(stack) - 1
		root = stack[top]
		stack = stack[:top]
	}
	return res
}

func preorder_morris_break(root *TreeNode) []int {
	var res []int
	for root != nil {
		res = append(res, root.Val)
		if root.Left == nil {
			root = root.Right
		} else {
			max := root.Left
			for max.Right != nil {
				max = max.Right
			}
			max.Right = root.Right
			root, root.Left = root.Left, nil
		}
	}
	return res
}

func preoder_morris_keep(root *TreeNode) []int {
	var res []int
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			max := root.Left
			for max.Right != nil && max.Right != root {
				max = max.Right
			}
			if max.Right == nil {
				res = append(res, root.Val)
				max.Right = root
				root = root.Left
			} else {
				root = root.Right
				max.Right = nil
			}
		}
	}
	return res
}
