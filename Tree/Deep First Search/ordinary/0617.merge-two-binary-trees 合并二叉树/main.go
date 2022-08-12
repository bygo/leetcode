package main

// https://leetcode.cn/problems/merge-two-binary-trees/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 合并二叉树

func mergeTrees(r1 *TreeNode, r2 *TreeNode) *TreeNode {
	if r1 == nil {
		return r2
	} else if r2 == nil {
		return r1
	}
	r1.Val = r1.Val + r2.Val
	r1.Left = mergeTrees(r1.Left, r2.Left)
	r1.Right = mergeTrees(r1.Right, r2.Right)
	return r1
}
