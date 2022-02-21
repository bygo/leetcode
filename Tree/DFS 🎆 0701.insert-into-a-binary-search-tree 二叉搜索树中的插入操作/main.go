package main

// https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉搜索树中的插入操作

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 寻路
	if root == nil {
		return nil
	}
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			// 进入左子树
			insertIntoBST(root.Left, val)
		}
	} else if root.Val < val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			// 进入右子树
			insertIntoBST(root.Right, val)
		}
	}

	return root
}
