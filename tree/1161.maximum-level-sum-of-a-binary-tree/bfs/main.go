package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Title: Maximum Level Sum of a Binary Tree
//Link: https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue = []*TreeNode{root}
	var res, level, sum int
	var max = -1 << 63
	for {
		counter := len(queue)
		if counter == 0 {
			break
		}
		level++
		for _, q := range queue[:counter] {
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
			res = level
		}
		sum = 0
		queue = queue[counter:]
	}

	return res
}
