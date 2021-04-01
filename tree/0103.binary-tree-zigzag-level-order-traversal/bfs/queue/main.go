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
	var level, counter, index, symbol int
	for {
		counter = len(queue)
		if counter == 0 {
			break
		}
		if level%2 == 0 {
			index = -1
			symbol = 1
		} else {
			index = counter
			symbol = -1
		}

		res = append(res, []int{})
		for _, q := range queue[:counter] {
			index += symbol
			res[level] = append(res[level], queue[index].Val)
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		level++
		queue = queue[counter:]
	}

	return res
}
