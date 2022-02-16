package main

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 从前序与中序遍历序列构造二叉树

func buildTree(preorder []int, inorder []int) *TreeNode {
	for idx := range inorder {
		if preorder[0] == inorder[idx] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:idx+1], inorder[0:idx]),
				Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
			}
		}
	}

	return nil
}
