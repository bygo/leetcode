package main

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// ❓ 填充每个节点的下一个右侧节点指针
// ⚠️ Perfect Binary Tree
// ⚠️ 左节点不断往右，右节点不断往左,等价于 x*2+1 连向 (x+1)*2

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	left := root.Left
	right := root.Right
	for left != nil {
		left.Next = right
		left = left.Right
		right = right.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}
