package main

// https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

// ❓ N 叉树的后序遍历

func postorder(root *Node) []int {
	var nums []int
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, child := range node.Children {
			dfs(child)
		}
		nums = append(nums, node.Val)
	}
	dfs(root)
	return nums
}
