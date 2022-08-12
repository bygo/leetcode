package main

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉搜索树的最近公共祖先

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			// p q 分别位于左右子树 或者当前节点
			return root
		}
	}
}
