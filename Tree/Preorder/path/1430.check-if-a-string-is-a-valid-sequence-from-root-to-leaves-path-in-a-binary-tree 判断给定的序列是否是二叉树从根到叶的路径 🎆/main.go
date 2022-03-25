package main

// https://leetcode-cn.com/problems/check-if-a-string-is-a-valid-sequence-from-root-to-leaves-path-in-a-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 判断给定的序列是否是二叉树从根到叶的路径

func isValidSequence(root *TreeNode, arr []int) bool {
	arrL := len(arr)
	var dfs func(node *TreeNode, idx int) bool
	dfs = func(node *TreeNode, idx int) bool {
		if node == nil {
			return false
		}

		if node.Val != arr[idx] {
			return false
		}
		idx++
		if idx == arrL {
			return node.Left == nil && node.Right == nil
		}

		return dfs(node.Left, idx) || dfs(node.Right, idx)
	}
	return dfs(root, 0)
}
