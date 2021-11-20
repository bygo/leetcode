package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// minimum depth of binary tree
// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

func minDepth(root *TreeNode) int {
	var depth int
	if root == nil {
		return depth
	}
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		depth++
		var cnt = len(queue)
		for _, q := range queue[:cnt] {
			if q.Left == nil && q.Right == nil {
				return depth
			}
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		queue = queue[cnt:]
	}
	return depth
}
