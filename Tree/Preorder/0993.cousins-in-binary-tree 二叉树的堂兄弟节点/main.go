package main

// https://leetcode-cn.com/problems/cousins-in-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的堂兄弟节点

func isCousins(root *TreeNode, x int, y int) bool {
	var pX, pY = -1, -1
	var depX, depY int
	var dfs func(node *TreeNode, p int, dep int)
	dfs = func(node *TreeNode, p int, dep int) {
		if node == nil {
			return
		}
		if node.Val == x {
			pX = p
			depX = dep
		} else if node.Val == y {
			pY = p
			depY = dep
		}
		if pX != -1 && pY != -1 {
			return
		}
		dfs(node.Left, node.Val, dep+1)
		dfs(node.Right, node.Val, dep+1)
	}
	dfs(root, -1, -1)
	return depX == depY && pX != pY
}
