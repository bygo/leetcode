package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	var level int
	var queue = []*Node{root}
	for {
		counter := len(queue)
		if counter == 0 {
			break
		}
		for _, v := range queue[:counter] {
			for _, c := range v.Children {
				queue = append(queue, c)
			}
		}
		queue = queue[counter:]
		level++
	}
	return level
}
