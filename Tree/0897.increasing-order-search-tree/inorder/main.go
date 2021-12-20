package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Title: Increasing Order Search Tree
// https://leetcode-cn.com/problems/increasing-order-search-tree

var last *TreeNode
var head *TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	last = nil
	head = nil
	dfs(root)
	return head
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		if last == nil {
			head = root
		} else {
			last.Right = root
		}

		root.Left = nil
		last = root
		dfs(root.Right)
	}
}
