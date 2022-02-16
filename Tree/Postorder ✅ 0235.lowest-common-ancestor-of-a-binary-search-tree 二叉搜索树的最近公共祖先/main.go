package main

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉搜索树的最近公共祖先

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node := root
	for {
		if p.Val < node.Val && q.Val < node.Val {
			node = node.Left
		} else if node.Val < p.Val && node.Val < q.Val {
			node = node.Right
		} else {
			return node
		}
	}
}
