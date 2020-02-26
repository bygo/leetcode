package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue = []*TreeNode{root}
	var minQ = []int{-1 << 63}
	var maxQ = []int{1<<63 - 1}
	for len(queue) > 0 {
		root, min, max := queue[0], minQ[0], maxQ[0]
		if root.Val <= min || max <= root.Val {
			return false
		}
		//push
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
		//shift
		queue, minQ, maxQ = queue[1:], minQ[1:], maxQ[1:]
	}

	return true
}
