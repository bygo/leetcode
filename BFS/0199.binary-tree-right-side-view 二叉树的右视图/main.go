package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// binary tree right side view
// https://leetcode-cn.com/problems/binary-tree-right-side-view/

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		var cnt = len(queue)
		for i, q := range queue[:cnt] {
			if i == cnt-1 {
				res = append(res, q.Val)
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
	return res
}
