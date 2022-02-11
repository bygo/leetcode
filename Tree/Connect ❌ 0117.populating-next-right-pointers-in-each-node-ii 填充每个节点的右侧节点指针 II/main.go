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

	var parent *Node
	parent = root

	for parent != nil {
		var pre, down *Node
		down = nil
		for parent != nil {
			if parent.Left != nil {
				if pre == nil {
					down = parent.Left
				} else {
					pre.Next = parent.Left
				}
				pre = parent.Left
			}

			if parent.Right != nil {
				if pre == nil {
					down = parent.Right
				} else {
					pre.Next = parent.Right
				}
				pre = parent.Right
			}
			parent = parent.Next
		}
		parent = down
	}
	return root
}
