package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}
	var level int
	for 0 < len(queue) {
		length := len(queue)
		for 0 < length {
			length--
			if len(res) == level {
				res = append(res, queue[0].Val)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}

			queue = queue[1:]
		}
		level++
	}
	return res
}
