package main

// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// ❓ 填充每个节点的下一个右侧节点指针

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var parent = root
	for parent.Left != nil {
		var p = parent
		for p != nil {
			p.Left.Next = p.Right
			if p.Next != nil {
				p.Right.Next = p.Next.Left
			}
			p = p.Next
		}
		parent = parent.Left
	}

	return root
}
