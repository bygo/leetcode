package main

// https://leetcode.cn/problems/n-ary-tree-preorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

// ❓ N 叉树的前序遍历

func preorder(root *Node) []int {
	var nums []int
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		nums = append(nums, node.Val)
		for _, child := range node.Children {
			dfs(child)
		}
	}
	dfs(root)
	return nums
}
