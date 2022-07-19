package main

// https://leetcode.cn/problems/delete-nodes-and-return-forest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 删点成林

func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	var numMp = map[int]bool{}
	for _, num := range toDelete {
		numMp[num] = true
	}

	// 根节点
	var nodes []*TreeNode
	if !numMp[root.Val] {
		nodes = append(nodes, root)
	}

	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			// 空节点不需要删除
			return false
		}
		bLeft := dfs(node.Left)
		bRight := dfs(node.Right)
		if bLeft {
			node.Left = nil
		}
		if bRight {
			node.Right = nil
		}

		// 不需要删除,也就没有拆分
		if !numMp[node.Val] {
			return false
		}

		// 拆分存储
		if node.Left != nil {
			nodes = append(nodes, node.Left)
			node.Left = nil
		}

		if node.Right != nil {
			nodes = append(nodes, node.Right)
			node.Right = nil
		}
		return true
	}
	dfs(root)

	return nodes
}
