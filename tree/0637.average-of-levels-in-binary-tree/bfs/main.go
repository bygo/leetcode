package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	var res []float64
	for {
		var sum int
		counter := len(queue)
		if counter == 0 {
			break
		}
		for _, q := range queue[:counter] {
			sum += q.Val
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		res = append(res, float64(sum)/float64(counter))
		queue = queue[counter:]
	}
	return res
}
