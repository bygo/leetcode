package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// populating next right pointers in each node
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var queue = []*Node{root}
	var head = root
	var pre *Node
	for 0 < len(queue) {
		var cnt = len(queue)
		pre = nil
		for _, v := range queue[:cnt] {
			if pre != nil {
				pre.Next = v
			}
			pre = v
			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		queue = queue[cnt:]
	}

	return head
}

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
