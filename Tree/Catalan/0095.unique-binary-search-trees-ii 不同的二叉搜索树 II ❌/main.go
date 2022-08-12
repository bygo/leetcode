package main

// https://leetcode.cn/problems/unique-binary-search-trees-ii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 不同的二叉搜索树 II

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	var dfs func(start, end int) []*TreeNode
	dfs = func(start, end int) []*TreeNode {
		var comb []*TreeNode
		if end < start {
			comb = append(comb, nil)
			return comb
		}
		for idx := start; idx <= end; idx++ {
			for _, left := range dfs(start, idx-1) {
				for _, right := range dfs(idx+1, end) {
					comb = append(comb, &TreeNode{
						Val:   idx,
						Left:  left,
						Right: right,
					})
				}
			}
		}
		return comb
	}
	return dfs(1, n)
}
