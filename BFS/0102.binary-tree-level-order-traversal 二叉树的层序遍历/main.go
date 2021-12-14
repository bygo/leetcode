package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var depth int
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		var cnt = len(queue)
		res = append(res, []int{})
		for _, q := range queue[:cnt] {
			res[depth] = append(res[depth], q.Val)
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		queue = queue[cnt:]
		depth++
	}

	return res
}
