package _0_11_12

func inorder(root *TreeNode) []int {
	var res []int
	idfs(root, res)
	return res
}

func idfs(root *TreeNode, res []int) {
	if root != nil {
		idfs(root.Left, res)
		res = append(res, root.Val)
		idfs(root.Right, res)
	}
}

func inorderStack(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack = []*TreeNode{root}
	for 0 < len(stack) {
		for root.Left != nil {
			stack = append(stack, root)
			root = root.Left
		}

		top := len(stack) - 1
		res = append(res, stack[top].Val)
		root = stack[top].Right
		stack = stack[:top]
	}
	return res
}

func inorderMorrisBreak(root *TreeNode) []int {
	var res []int
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			max := root.Left
			for max.Right != nil {
				max = max.Right
			}

			max.Right = root
			root, root.Left = root.Left, nil
		}
	}
	return res
}

func inorderMorrisKeep(root *TreeNode) []int {
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
				max.Right = root
				root = root.Left
			} else {
				res = append(res, root.Val)
				root = root.Right
				max.Right = nil
			}
		}
	}
	return res
}
