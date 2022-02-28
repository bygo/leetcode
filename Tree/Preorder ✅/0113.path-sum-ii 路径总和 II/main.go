package main

// https://leetcode-cn.com/problems/path-sum-ii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 路径总和 II

func pathSum(root *TreeNode, sum int) [][]int {
	var pathsSum [][]int
	var path []int
	var sumCur int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		sumCur += node.Val
		if node.Left == nil && node.Right == nil && sumCur == sum {
			pathsSum = append(pathsSum, append([]int{}, path...))
		}
		dfs(node.Left)
		dfs(node.Right)
		sumCur -= node.Val
		path = path[:len(path)-1]
	}
	dfs(root)
	return pathsSum
}
