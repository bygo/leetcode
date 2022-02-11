package main

// https://leetcode-cn.com/problems/find-nearest-right-node-in-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 找到二叉树中最近的右侧节点

func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var que = []*TreeNode{root}

	for {
		cnt := len(que)
		if cnt == 0 {
			break
		}

		for idx, q := range que[:cnt] {
			if q == u {
				if idx == cnt-1 {
					return nil
				} else {
					return que[idx+1]
				}
			}
			if q.Left != nil {
				que = append(que, q.Left)
			}
			if q.Right != nil {
				que = append(que, q.Right)
			}
		}
		que = que[cnt:]
	}
	return nil
}
