package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// Link: https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list

// 递归
func flatten(root *Node) *Node {
	cur := root
	for cur != nil {
		if cur.Child == nil {
			cur = cur.Next // 移动
		} else {
			next := cur.Next        // 临时保存下一节点
			child := cur.Child      // 临时保存子节点
			cur.Next = child        // 指向子节点 head
			cur.Child = nil         // 解关系
			child.Prev = cur        // 前置指向
			cHead := flatten(child) // 递归处理
			for cHead.Next != nil {
				cHead = cHead.Next // 移动到尾节点
			}

			if next != nil {
				next.Prev = cHead // 下一节点不为空 前置指向
			}
			cHead.Next = next // 后置指向
			cur = next        // 移动
		}
	}
	return root
}

// 迭代
func flatten2(root *Node) *Node {
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
