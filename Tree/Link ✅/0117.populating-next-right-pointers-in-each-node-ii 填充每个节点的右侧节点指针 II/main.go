package main

// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// ❓ 填充每个节点的下一个右侧节点指针 II

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var parent, child, link *Node
	parent = root

	suture := func(node *Node) {
		if node == nil {
			return
		}
		if link == nil {
			// 头指针
			child = node
		} else {
			link.Next = node
		}
		link = node
	}

	for parent != nil {
		link, child = nil, nil
		// 顶层链表移动,连接底层链表
		for parent != nil {
			suture(parent.Left)
			suture(parent.Right)
			parent = parent.Next
		}
		parent = child
	}
	return root
}
