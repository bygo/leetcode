package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// average of levels in binary tree
// https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	var res []float64
	for 0 < len(queue) {
		var sum int
		var cnt = len(queue)
		for _, q := range queue[:cnt] {
			sum += q.Val
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		res = append(res, float64(sum)/float64(cnt))
		queue = queue[cnt:]
	}
	return res
}
