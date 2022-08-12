package main

// https://leetcode.cn/problems/find-leaves-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 寻找二叉树的叶子节点

func findLeaves(root *TreeNode) [][]int {
	var depsNums [][]int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}

		depLeft := dfs(node.Left)
		depRight := dfs(node.Right)
		dep := max(depLeft, depRight) + 1 // 自底向上搜集
		if len(depsNums) <= dep {
			depsNums = append(depsNums, []int{})
		}
		depsNums[dep] = append(depsNums[dep], node.Val)
		return dep
	}
	dfs(root)
	return depsNums
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
