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
	pre := root

	for pre.Left != nil {
		parent := pre
		for parent != nil {
			parent.Left.Next = parent.Right
			if parent.Next != nil {
				parent.Right.Next = parent.Next.Left
			}
			parent = parent.Next
		}
		pre = pre.Left
	}
	return root
}
