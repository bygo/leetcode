package main

// https://leetcode.cn/problems/binary-tree-right-side-view/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的右视图

func rightSideView(root *TreeNode) []int {
	var nums []int
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
			if idx == queL-1 {
				nums = append(nums, node.Val)
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
	return nums
}
