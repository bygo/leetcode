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
			parent.Left.Next = parent.Right //左节点连接右节点
			if parent.Next != nil {
				parent.Right.Next = parent.Next.Left //右节点 连接 邻居左节点
			}
			parent = parent.Next //同层移动
		}
		pre = pre.Left //移到下层
	}
	return root
}
