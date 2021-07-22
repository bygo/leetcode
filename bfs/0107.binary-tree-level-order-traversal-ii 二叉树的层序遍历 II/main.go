package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// binary tree level order traversal (bottom)
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue = []*TreeNode{root}
	var depth int
	for 0 < len(queue) {
		var cnt = len(queue)
		res = append(res, []int{})
		for _, q := range queue[:cnt] {
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, q.Right)
			}
			res[depth] = append(res[depth], q.Val)
		}
		queue = queue[cnt:]
		depth++
	}

	var i = 0
	var j = len(res) - 1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}
