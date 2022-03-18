package main

// https://leetcode-cn.com/problems/split-bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 拆分二叉搜索树
func splitBST(root *TreeNode, target int) []*TreeNode {
	if root == nil {
		return []*TreeNode{nil, nil}
	} else if root.Val <= target {
		// 左树合法，进入右树

		ts := splitBST(root.Right, target)
		root.Right = ts[0] // right 连上合法位置
		ts[0] = root       // 设当前节点为第一位置
		return ts
	} else {
		// 右树合法，进入左树
		ts := splitBST(root.Left, target)
		root.Left = ts[1] // left 连上合法位置
		ts[1] = root      // 设当前节点为第二位置
		return ts
	}
}
