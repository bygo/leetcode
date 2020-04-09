package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var queue = []*Node{root}
	var level int
	for 0 < len(queue) {
		length := len(queue)
		for 0 < length {
			length--
			for _, n := range queue[0].Children {
				queue = append(queue, n)
			}
			queue = queue[1:]
		}
		level++
	}
	return level
}
