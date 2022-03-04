package main

// https://leetcode-cn.com/problems/equal-tree-partition

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkEqualTree(root *TreeNode) bool {
	sumMp := map[int]int{}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := dfs(node.Left) + dfs(node.Right) + node.Val
		sumMp[sum] ++
		return sum
	}

	tot := dfs(root)
	if tot&1 == 1 {
		return false
	}
	sumMp[tot]--
	return 0 < sumMp[tot/2]
}
