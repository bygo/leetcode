package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Link: https://leetcode-cn.com/problems/reverse-linked-list-ii

func reverseBetween(head *ListNode, left, right int) *ListNode {
	zero := &ListNode{Next: head}
	prev := zero
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	//第一个不管 从下一个开始插入
	p := prev.Next
	for i := left; i < right; i++ {
		first := p.Next        // 等待插入
		p.Next = p.Next.Next   // 删除指针
		first.Next = prev.Next // 插入 head
		prev.Next = first      // 插入 完成
	}
	return zero.Next
}

func reverse(head *ListNode) *ListNode {
	if head != nil {
		cur := head.Next
		head.Next = nil
		for cur != nil {
			next := cur.Next
			cur.Next = head
			head = cur
			cur = next
		}
	}
	return head
}
