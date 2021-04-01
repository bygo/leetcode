package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var level int
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}
	for {
		counter := len(queue)
		if counter == 0 {
			break
		}
		res = append(res, []int{})
		for _, q := range queue[:counter] {
			res[level] = append(res[level], q.Val)
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		queue = queue[counter:]
		level++
	}

	return res
}
