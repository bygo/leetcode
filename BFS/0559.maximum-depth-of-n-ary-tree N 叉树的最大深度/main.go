package main

type Node struct {
	Val      int
	Children []*Node
}

// maximum depth of n-ary-tree
// https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/submissions/

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	var depth int
	var queue = []*Node{root}
	for 0 < len(queue) {
		var cnt = len(queue)
		for _, q := range queue[:cnt] {
			for _, c := range q.Children {
				queue = append(queue, c)
			}
		}
		queue = queue[cnt:]
		depth++
	}
	return depth
}
