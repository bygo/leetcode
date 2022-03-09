package main

// https://leetcode-cn.com/problems/maximum-binary-tree-ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 最大二叉树 II
// ⚠️ 因为添加在末端，添加的值一定在右子树

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val < val {
		return &TreeNode{Left: root, Val: val}
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}
