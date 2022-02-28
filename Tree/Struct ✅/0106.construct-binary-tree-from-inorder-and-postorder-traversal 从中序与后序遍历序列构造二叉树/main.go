package main

// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	numsL := len(postorder) - 1
	for idx := range inorder {
		if inorder[idx] == postorder[numsL] {
			return &TreeNode{
				Val:   postorder[numsL],
				Left:  buildTree(inorder[:idx], postorder[0:idx]),
				Right: buildTree(inorder[idx+1:], postorder[idx:numsL]),
			}
		}
	}
	return nil
}
