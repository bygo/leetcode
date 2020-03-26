package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	var level int
	if root == nil {
		return level
	}
	var queue = []*TreeNode{root}

	for len(queue) > 0 {
		level++
		length := len(queue)
		for 0 < length {
			length--
			if queue[0].Left == nil && queue[0].Right == nil {
				return level
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}

			queue = queue[1:]
		}
	}
	return level
}
