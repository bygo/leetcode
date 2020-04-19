package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: v}
	}
	if  d == 1 {
		return &TreeNode{Left: root,Val:v}
	}
	var queue = []*TreeNode{root}
	var depth = 1
	for 0 < len(queue) {
		depth++
		var l = len(queue)
		if depth == d {
			for i := 0; i < l; i++ {
				nL := &TreeNode{Left: queue[i].Left, Val: v}
				queue[i].Left = nL
				nR := &TreeNode{Right: queue[i].Right, Val: v}
				queue[i].Right = nR
			}
		}
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}
	return root
}
