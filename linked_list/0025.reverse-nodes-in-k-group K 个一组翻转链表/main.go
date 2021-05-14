package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Link: https://leetcode-cn.com/problems/reverse-nodes-in-k-group

func reverseKGroup(head *ListNode, k int) *ListNode {
	zero := &ListNode{Next: head}
	prev := zero
	for head != nil {
		tail := head
		for i := 1; i < k && tail != nil; i++ { //找链表末尾
			tail = tail.Next
		}
		if tail == nil { //不够返回
			break
		}

		next := tail.Next              // 保存末尾下一个
		tail.Next = nil                // 置空末尾下一个
		prev.Next = reverse(prev.Next) // 反转链表
		head.Next = next               // 头变尾， 末尾 连接 下一个
		prev = head                    // 前置指针
		head = head.Next               // 下一个链表
	}
	return zero.Next
}

//1->2->3
func reverse(head *ListNode) *ListNode {
	if head != nil {
		p := head.Next
		head.Next = nil
		for p != nil {
			next := p.Next // 保存下一个
			p.Next = head  // 2 指向 1
			head = p       // 1 移到 2
			p = next       // 2 移到 3
		}
	}
	return head
}
