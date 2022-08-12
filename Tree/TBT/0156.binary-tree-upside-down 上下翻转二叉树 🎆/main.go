package main

// https://leetcode.cn/problems/binary-tree-upside-down/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 上下翻转二叉树
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	var pre, left, right *TreeNode

	for root != nil {
		left = root.Left // 下一次循环

		root.Left = right  // ⚠️ 上一次右节点 变成 当前节点的左节点
		right = root.Right // 下一次循环

		root.Right = pre // ⚠️ 上一次根节点 变成 当前节点的右节点
		pre = root       // 下一次循环
		root = left      // 下一次循环
	}
	return pre
}
