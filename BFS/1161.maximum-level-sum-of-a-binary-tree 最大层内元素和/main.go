package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maximum level sum of a binary tree
// https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue = []*TreeNode{root}
	var res, depth, sum int
	var max = -1 << 63
	for 0 < len(queue) {
		var cnt = len(queue)
		depth++
		for _, q := range queue[:cnt] {
			sum += q.Val
			if q.Left != nil {
				queue = append(queue, q.Left)
			}

			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		if max < sum {
			max = sum
			res = depth
		}
		sum = 0
		queue = queue[cnt:]
	}

	return res
}
