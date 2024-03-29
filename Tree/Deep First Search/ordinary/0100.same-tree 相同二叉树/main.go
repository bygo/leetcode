package main

// https://leetcode.cn/problems/same-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 相同二叉树

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return isSameTree(p.Left, q.Left) && p.Val == q.Val && isSameTree(p.Right, q.Right)
}
