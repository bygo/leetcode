package main

// https://leetcode.cn/problems/increasing-order-search-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 递增顺序搜索树

func increasingBST(root *TreeNode) *TreeNode {
	nodeZero := &TreeNode{}
	nodeCur := nodeZero

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		nodeCur.Right = node
		node.Left = nil
		nodeCur = nodeCur.Right
		dfs(node.Right)
	}
	dfs(root)

	return nodeZero.Right
}
