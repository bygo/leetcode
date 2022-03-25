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
	var getDepMax func(node *TreeNode) int
	getDepMax = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		depL := getDepMax(node.Left)
		depR := getDepMax(node.Right)
		return max(depL, depR) + 1
	}
	depMax = getDepMax(root)

	var depsStrs = make([][]string, depMax)

	// 预处理
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
		depsStrs[dep][idx] = strconv.Itoa(node.Val)
		dep += 1
		// 每次half>>dep 都舍1，需补上
		dfs(node.Left, idx-half>>dep-1, dep)
		dfs(node.Right, idx+half>>dep+1, dep)
	}
	dfs(root, half, 0)
	return depsStrs
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
