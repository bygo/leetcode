package main

// https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 把二叉搜索树转换为累加树
// ⚠️ BST
// ⚠️ 节点 = sum(所有大于当前节点)

func bstToGst(root *TreeNode) *TreeNode {
	var sum int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}

// 0538 https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
