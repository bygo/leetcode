package main

// https://leetcode-cn.com/problems/split-bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 拆分二叉搜索树

func splitBST(root *TreeNode, target int) []*TreeNode {
	var dfs func(node *TreeNode) (*TreeNode, *TreeNode)
	dfs = func(node *TreeNode) (*TreeNode, *TreeNode) {
		if node == nil {
			return nil, nil
		}

		if node.Val <= target {
			// 左树合法，进入右树
			left, right := dfs(node.Right)
			node.Right = left // right 连上合法位置
			return node, right
		} else {
			// 右树合法，进入左树
			left, right := dfs(node.Left)
			node.Left = right // left 连上合法位置
			return left, node
		}
	}

	left, right := dfs(root)
	return []*TreeNode{left, right}
}
