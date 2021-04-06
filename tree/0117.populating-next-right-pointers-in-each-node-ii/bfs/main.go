package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var cur *Node
	cur = root

	for cur != nil {
		var pre, down *Node
		down = nil
		for cur != nil {
			if cur.Left != nil {
				if pre == nil {
					down = cur.Left
				} else {
					pre.Next = cur.Left
				}
				pre = cur.Left
			}

			if cur.Right != nil {
				if pre == nil {
					down = cur.Right
				} else {
					pre.Next = cur.Right
				}
				pre = cur.Right
			}
			cur = cur.Next
		}
		cur = down
	}
	return root
}
