package main

// https://leetcode-cn.com/problems/insufficient-nodes-in-root-to-leaf-paths

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 根到叶路径上的不足节点

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	var dfs func(node *TreeNode, numCur int) bool
	dfs = func(node *TreeNode, numCur int) bool {
		if node == nil {
			return true
		}

		numCur += node.Val
		if node.Left == nil && node.Right == nil {
			return numCur < limit
		}
		bLeft := dfs(node.Left, numCur)
		bRight := dfs(node.Right, numCur)
		if bLeft {
			node.Left = nil
		}
		if bRight {
			node.Right = nil
		}
		return bLeft && bRight
	}
	if dfs(root, 0) {
		return nil
	}

	return root
}
