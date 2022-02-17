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
		for _, node := range que[:queL] {
			if node.Left == nil && node.Right == nil {
				return dep
			}
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		que = que[queL:]
	}
	return dep
}
