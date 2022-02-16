package main

// https://leetcode-cn.com/problems/inorder-successor-in-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉搜索树中的中序后继

// https://github.com/bygo/leetcode

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var stack []*TreeNode
	var pre *TreeNode
	for 0 < len(stack) || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		top := len(stack) - 1
		if pre != nil {
			return stack[top]
		} else if stack[top] == p {
			pre = p
		}
		root = stack[top].Right
		stack = stack[:top]
	}
	return nil
}
