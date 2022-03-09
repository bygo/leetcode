package main

import "strconv"

// https://leetcode-cn.com/problems/print-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 输出二叉树

func printTree(root *TreeNode) [][]string {
	var depMax int
	var getDep func(node *TreeNode) int
	getDep = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		depL := getDep(node.Left)
		depR := getDep(node.Right)
		return max(depL, depR) + 1
	}
	depMax = getDep(root)

	var depsStrs = make([][]string, depMax)

	bufL := 1<<depMax - 1
	for dep := 0; dep < depMax; dep++ {
		depsStrs[dep] = make([]string, bufL)
		for idx := 0; idx < bufL; idx++ {
			depsStrs[dep][idx] = ""
		}
	}

	half := bufL / 2
	var dfs func(node *TreeNode, idx int, dep int)
	dfs = func(node *TreeNode, idx int, dep int) {
		if node == nil {
			return
		}
		depsStrs[dep][idx<<(depMax-dep-1)+half] = strconv.Itoa(node.Val)
		dfs(node.Left, idx<<1-1, dep+1)
		dfs(node.Right, idx<<1+1, dep+1)

	}
	dfs(root, 0, 0)
	return depsStrs
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
