package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// validate binary search tree is valid ?
// https://leetcode-cn.com/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var queue = []*TreeNode{root}
	var minQ = []int{-1 << 63}
	var maxQ = []int{1<<63 - 1}
	for 0 < len(queue) {
		var cnt = len(queue)
		for i, root := range queue[:cnt] {
			min, max := minQ[i], maxQ[i]
			if root.Val <= min || max <= root.Val {
				return false
			}
			if root.Left != nil {
				minQ = append(minQ, min)
				maxQ = append(maxQ, root.Val)
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				minQ = append(minQ, root.Val)
				maxQ = append(maxQ, max)
				queue = append(queue, root.Right)
			}
		}
		queue, minQ, maxQ = queue[cnt:], minQ[cnt:], maxQ[cnt:]
	}

	return true
}
