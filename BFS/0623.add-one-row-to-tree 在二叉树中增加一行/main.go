package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// add one row to tree
// https://leetcode-cn.com/problems/add-one-row-to-tree/

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: v}
	}

	if d == 1 {
		return &TreeNode{Val: v, Left: root}
	}

	var queue = []*TreeNode{root}
	var depth = 1

	for 0 < len(queue) {
		var cnt = len(queue)
		depth++
		if depth == d {
			for _, q := range queue {
				nl := &TreeNode{Val: v, Left: q.Left}
				q.Left = nl
				nr := &TreeNode{Val: v, Right: q.Right}
				q.Right = nr
			}
			break
		}

		for _, q := range queue[:cnt] {
			if q.Left != nil {
				queue = append(queue, q.Left)
			}

			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		queue = queue[cnt:]
	}

	return root
}
