package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(node *TreeNode) bool {
	return find(node) > -1
}

func find(node *TreeNode) float64 {
	if node == nil {
		return 0
	}

	l := find(node.Left)
	if l == -1 { //剪枝，不平衡时直接返回
		return -1
	}
	r := find(node.Right)

	if l == -1 || r == -1 || 1 < math.Abs(l-r) { //剪枝，不平衡时直接返回
		return -1
	}

	return math.Max(l, r) + 1 //计算左右最大深度，因为是后序遍历，游标会一直走到最底部，再往上冒泡
}
