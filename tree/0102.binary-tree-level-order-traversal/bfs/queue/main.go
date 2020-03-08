package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue)
		res[level] = make([]int, 0)
		for 0 < counter {
			counter--
			if queue[counter].Left != nil {
				queue = append(queue, queue[counter].Left)
			}
			if queue[counter].Right != nil {
				queue = append(queue, queue[counter].Right)
			}
			res[level] = append(res[level], queue[counter].Val)
		}
		level++
	}
	return res
}
