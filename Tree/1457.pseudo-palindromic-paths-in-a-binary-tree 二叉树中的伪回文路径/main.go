package main

// https://leetcode.cn/problems/pseudo-palindromic-paths-in-a-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树中的伪回文路径

func pseudoPalindromicPaths(root *TreeNode) int {
	cnt := 0
	var dfs func(node *TreeNode, subset int)
	dfs = func(node *TreeNode, subset int) {
		if node == nil {
			return
		}
		subset ^= 1 << node.Val
		if node.Left == nil && node.Right == nil {
			if subset == 0 || subset&(subset-1) == 0 {
				cnt++
			}
			return
		}
		dfs(node.Left, subset)
		dfs(node.Right, subset)

	}
	dfs(root, 0)

	return cnt
}
