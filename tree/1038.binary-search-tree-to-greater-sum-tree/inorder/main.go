package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Title: Binary Search Tree to Greater Sum Tree
// Link: https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree

func bstToGst(root *TreeNode) *TreeNode {
	var sum int
	dfs(root, &sum)
	return root
}

func dfs(root *TreeNode, sum *int) {
	if root != nil {
		dfs(root.Right, sum)
		*sum += root.Val
		root.Val = *sum
		dfs(root.Left, sum)
	}
}
