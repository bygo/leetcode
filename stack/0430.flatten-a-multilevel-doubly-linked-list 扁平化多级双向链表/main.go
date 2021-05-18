package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

//Link: https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	zero := &Node{Next: root}
	prev := zero
	stack := []*Node{root}

	for {
		top := len(stack) - 1
		if top < 0 {
			break
		}

		cur := stack[top]
		stack = stack[:top]

		prev.Next = cur
		cur.Prev = prev

		if cur.Next != nil {
			stack = append(stack, cur.Next)
		}

		if cur.Child != nil {
			stack = append(stack, cur.Child)
			cur.Child = nil
		}
		prev = cur
	}

	zero.Next.Prev = nil

	return zero.Next
}
