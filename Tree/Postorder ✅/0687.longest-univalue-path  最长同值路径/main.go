package main

// https://leetcode-cn.com/problems/longest-univalue-path/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 最长同值路径

func longestUnivaluePath(root *TreeNode) int {
	depMax := 0
	var dfs func(node *TreeNode, valParent int, dep int) int
	dfs = func(node *TreeNode, valParent int, dep int) int {
		if node == nil {
			return 0
		}
		pathLeft := dfs(node.Left, node.Val, dep+1)   // 左树最长路径
		pathRight := dfs(node.Right, node.Val, dep+1) // 右树最长路径
		depMax = max(depMax, pathRight+pathLeft)      // 路径是否变长
		if valParent != node.Val {                    // 不相等时重新计算
			return 0
		}
		return max(pathLeft, pathRight) + 1 // 相等时+1，并且选择l或r中最长路径
	}
	dfs(root, 0, 0)
	return depMax
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
