package main

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的最近公共祖先

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		// 最终返回 nil p q
		return root
	}
	nodeLeft := lowestCommonAncestor(root.Left, p, q)
	nodeRight := lowestCommonAncestor(root.Right, p, q)
	if nodeLeft == nil {
		return nodeRight
	}
	if nodeRight == nil {
		return nodeLeft
	}

	// 都不为空 ，本节点为祖先节点
	return root
}
