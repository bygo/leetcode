package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	var max *Node
	for root != nil {
		if root.Left == nil {
			root = root.Right
		} else {
			max = root.Left
			for max.Right != nil {
				max = max.Right
			}


		}
	}
}
