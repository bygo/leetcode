package main

// https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的完全性检验

func isCompleteTree(root *TreeNode) bool {
	empty := false
	que := []*TreeNode{root}
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		for _, node := range que {
			if node != nil {
				if empty {
					return false
				}
				que = append(que, node.Left)
				que = append(que, node.Right)
			} else {
				empty = true
			}
		}
		que = que[queL:]
	}
	return true
}
