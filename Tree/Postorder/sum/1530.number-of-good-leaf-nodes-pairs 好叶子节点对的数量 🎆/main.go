package main

// https://leetcode-cn.com/problems/number-of-good-leaf-nodes-pairs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 好叶子节点对的数量

func countPairs(root *TreeNode, distance int) int {
	var cntRes int
	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{}
		}
		if node.Left == nil && node.Right == nil {
			return []int{0}
		}
		depsLeft := dfs(node.Left)
		for i := range depsLeft {
			depsLeft[i]++
		}
		depsRight := dfs(node.Right)
		for i := range depsRight {
			depsRight[i]++
		}
		// 依次判断
		for _, depLeft := range depsLeft {
			for _, depRight := range depsRight {
				if depLeft+depRight <= distance {
					cntRes++
				}
			}
		}
		return append(depsLeft, depsRight...)
	}
	dfs(root)
	return cntRes
}
