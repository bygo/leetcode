package main

// https://leetcode-cn.com/problems/longest-univalue-path/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 最长同值路径

func longestUnivaluePath(root *TreeNode) int {
	pathMax := 0
	var dfs func(node *TreeNode, valParent int) int
	dfs = func(node *TreeNode, valParent int) int {
		if node == nil {
			return 0
		}
		pathLeft := dfs(node.Left, node.Val)       // 左树最长路径
		pathRight := dfs(node.Right, node.Val)     // 右树最长路径
		pathMax = max(pathMax, pathRight+pathLeft) // 路径是否变长
		if valParent != node.Val {                 // 不相等时重新计算
			return 0
		}
		return max(pathLeft, pathRight) + 1 // 相等时+1，并且选择l或r中最长路径
	}
	dfs(root, 0)
	return pathMax
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
