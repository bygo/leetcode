package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// find bottom left tree value
// https://leetcode-cn.com/problems/find-bottom-left-tree-value/

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int

	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		var cnt = len(queue)
		res = queue[0].Val
		for _, v := range queue[:cnt] {
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		queue = queue[cnt:]
	}
	return res
}
