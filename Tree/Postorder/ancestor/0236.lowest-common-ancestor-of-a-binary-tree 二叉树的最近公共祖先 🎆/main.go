package main

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的最近公共祖先

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		// 最终返回 nil
		return nil
	}
	if root == p || root == q {
		// p q ,可能 p或者q 为彼此祖先节点，到达根节点时，也只有一个p或者q
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
