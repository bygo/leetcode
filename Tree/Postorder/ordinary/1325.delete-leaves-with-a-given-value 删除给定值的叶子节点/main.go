package main

// https://leetcode.cn/problems/delete-leaves-with-a-given-value

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 删除给定值的叶子节点

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var remove func(node *TreeNode) *TreeNode
	remove = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		node.Left = remove(node.Left)
		node.Right = remove(node.Right)
		if node.Left == nil && node.Right == nil && node.Val == target {
			return nil
		}
		return node
	}
	return remove(root)
}
