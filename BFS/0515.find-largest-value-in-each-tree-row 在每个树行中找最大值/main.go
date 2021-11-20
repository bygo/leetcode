package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// find largest value in each tree row
// https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		var cnt = len(queue)
		var max = -1 << 63
		for _, v := range queue[:cnt] {
			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}
			if max < v.Val {
				max = v.Val
			}
		}
		res = append(res, max)
		queue = queue[cnt:]
	}

	return res
}
