package main

// https://leetcode.cn/problems/flip-equivalent-binary-trees/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 翻转等价二叉树

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root1 == root2
	}
	if root1.Val != root2.Val {
		return false
	}

	// 不翻转
	if flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) {
		return true
	}

	// 翻转
	if flipEquiv(root1.Left, root2.Right) && flipEquiv(root2.Left, root1.Right) {
		return true
	}
	return false
}
