package main

// https://leetcode.cn/problems/validate-binary-tree-nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 验证二叉树

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	var degree = make([]int, n)
	for _, child := range leftChild {
		if child == -1 {
			continue
		}
		degree[child]++
	}

	for _, child := range rightChild {
		if child == -1 {
			continue
		}
		degree[child]++
	}

	var root = -1
	var cntZero int
	for node, cnt := range degree {
		// 两个父节点
		if 1 < cnt {
			return false
		} else if cnt == 0 {
			// 两个根节点
			if cntZero == 1 {
				return false
			} else {
				cntZero++
				root = node
			}
		}
	}
	if root == -1 {
		return false
	}

	// 连通性
	var vis = map[int]struct{}{}
	var dfs func(node int)
	dfs = func(node int) {
		if node == -1 {
			return
		}
		vis[node] = struct{}{}
		dfs(leftChild[node])
		dfs(rightChild[node])
	}
	dfs(root)
	return len(vis) == n
}
