package main

// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的最小深度 

func minDepth(root *TreeNode) int {
	var dep int
	if root == nil {
		return dep
	}
	var que = []*TreeNode{root}
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		dep++
		for _, q := range que[:queL] {
			if q.Left == nil && q.Right == nil {
				return dep
			}
			if q.Left != nil {
				que = append(que, q.Left)
			}
			if q.Right != nil {
				que = append(que, q.Right)
			}
		}
		que = que[queL:]
	}
	return dep
}
