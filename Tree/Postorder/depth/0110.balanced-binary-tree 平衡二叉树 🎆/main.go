package main

// https://leetcode-cn.com/problems/balanced-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 平衡二叉树
// ⚠️ 二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1

func isBalanced(node *TreeNode) bool {
	return -1 < dfs(node)
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := dfs(node.Left)
	if left == -1 {
		return -1
	}

	right := dfs(node.Right)
	if right == -1 {
		return -1
	}

	if 1 < abs(left-right) { // 左右相差大于1
		return -1
	}

	return max(left, right) + 1 //计算左右最大深度
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
