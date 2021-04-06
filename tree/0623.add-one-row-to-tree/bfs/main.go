package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: v}
	}

	if d == 1 {
		return &TreeNode{Left: root, Val: v}
	}

	var queue = []*TreeNode{root}
	var depth = 1
	for {
		counter := len(queue)
		if counter == 0 {
			break
		}
		depth++

		if depth == d {
			for _, q := range queue[:counter] {
				newL := &TreeNode{Left: q.Left, Val: v}
				q.Left = newL
				newR := &TreeNode{Right: q.Right, Val: v}
				q.Right = newR
			}
			break
		}
		for _, q := range queue[:counter] {
			if q.Left != nil {
				queue = append(queue, q.Left)
			}

			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		queue = queue[counter:]
	}
	return root
}
