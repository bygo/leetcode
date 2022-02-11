package main

// https://leetcode-cn.com/problems/binary-tree-right-side-view/

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
		for idx, q := range que[:queL] {
			if idx == queL-1 {
				nums = append(nums, q.Val)
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
	return nums
}
