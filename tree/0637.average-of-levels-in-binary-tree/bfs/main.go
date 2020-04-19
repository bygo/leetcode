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
	for 0 < len(queue) {
		var l = len(queue)
		var counter int
		for i := 0; i < l; i++ {
			counter += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, float64(counter)/float64(l))
		queue = queue[l:]
	}
	return res
}
