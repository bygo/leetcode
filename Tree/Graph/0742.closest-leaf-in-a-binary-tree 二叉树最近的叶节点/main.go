package main

// https://leetcode.cn/problems/closest-leaf-in-a-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树最近的叶节点

func findClosestLeaf(root *TreeNode, k int) int {
	var graph = map[*TreeNode][]*TreeNode{}
	var que []*TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left != nil {
			graph[node] = append(graph[node], node.Left)
			graph[node.Left] = append(graph[node.Left], node)
			dfs(node.Left)
		}
		if node.Right != nil {
			graph[node] = append(graph[node], node.Right)
			graph[node.Right] = append(graph[node.Right], node)
			dfs(node.Right)
		}
		if node.Val == k {
			que = append(que, node)
		}
	}
	dfs(root)

	var depRes = 1<<63 - 1
	var vis = map[*TreeNode]bool{}
	var dep int
	var nodeRes int
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		dep++
		for _, node := range que[:queL] {
			vis[node] = true
			if node.Left == nil && node.Right == nil && dep < depRes {
				depRes = dep
				nodeRes = node.Val
			}
			for _, child := range graph[node] {
				if !vis[child] {
					que = append(que, child)
				}
			}
		}
		que = que[queL:]
	}

	return nodeRes
}
