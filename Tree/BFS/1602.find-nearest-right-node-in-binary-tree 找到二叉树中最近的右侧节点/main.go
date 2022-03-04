package main

// https://leetcode-cn.com/problems/find-nearest-right-node-in-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 找到二叉树中最近的右侧节点

func findNearestRightNode(root *TreeNode, target *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var que = []*TreeNode{root}

	for {
		queL := len(que)
		if queL == 0 {
			break
		}

		for idx, node := range que[:queL] {
			if node == target {
				if idx == queL-1 {
					return nil
				} else {
					return que[idx+1]
				}
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
	return nil
}
