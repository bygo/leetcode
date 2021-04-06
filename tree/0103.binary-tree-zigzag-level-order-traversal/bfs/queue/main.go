package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var queue = []*TreeNode{root}
	var level, counter int
	for {
		counter = len(queue)
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
		if level%2 == 1 {
			i := 0
			j := len(res[level]) - 1
			for i < j {
				res[level][i], res[level][j] = res[level][j], res[level][i]
				i++
				j--
			}
		}
		level++
		queue = queue[counter:]
	}

	return res
}
