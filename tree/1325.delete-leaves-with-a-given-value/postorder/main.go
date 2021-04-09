package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Title: Delete Leaves With a Given Value
//Link: https://leetcode-cn.com/problems/delete-leaves-with-a-given-value

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
