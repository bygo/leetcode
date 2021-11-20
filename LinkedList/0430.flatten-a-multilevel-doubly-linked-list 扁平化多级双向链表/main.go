package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list

// 递归

func flatten(root *Node) *Node {
	var dfs func(root *Node) *Node
	dfs = func(root *Node) *Node {
		var last *Node
		n := root
		for n != nil {
			next := n.Next
			if n.Child != nil {
				tail := dfs(n.Child)

				n.Next = n.Child
				n.Child.Prev = n

				if next != nil {
					tail.Next = next
					next.Prev = tail
				}

				n.Child = nil
				last = tail
			} else {
				last = n
			}
			n = next
		}
		return last
	}
	dfs(root)
	return root
}

// 迭代
func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	head := root
	prev := head
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

	head.Prev = nil

	return head
}
