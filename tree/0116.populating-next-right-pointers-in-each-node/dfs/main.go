package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//左节点不断往右，右节点不断往左，像拉链一样拉紧！
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	l := root.Left
	r := root.Right
	if l != nil {
		l.Next = r
		l = l.Right
		r = r.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}
