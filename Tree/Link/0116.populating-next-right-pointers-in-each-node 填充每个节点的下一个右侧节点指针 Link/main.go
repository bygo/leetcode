package main

// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// ❓ 填充每个节点的下一个右侧节点指针
// ⚠️ Complete Binary Tree

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var parent = root
	for parent.Left != nil {
		var link = parent
		for link != nil {
			// 本节点
			link.Left.Next = link.Right

			// 邻居
			if link.Next != nil {
				link.Right.Next = link.Next.Left
			}
			link = link.Next
		}
		parent = parent.Left
	}

	return root
}
