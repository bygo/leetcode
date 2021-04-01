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
	for {
		counter := len(queue)
		if counter == 0 {
			break
		}
		for _, n := range queue[:counter] {
			for _, v := range n.Children {
				queue = append(queue, v)
			}
		}
		queue = queue[counter:]
		level++
	}
	return level
}
