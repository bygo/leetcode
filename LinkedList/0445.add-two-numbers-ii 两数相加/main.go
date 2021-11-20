package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/add-two-numbers-ii

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)

	l := l1
	var carry = 0
	for {
		l1.Val = l1.Val + l2.Val + carry
		carry = l1.Val / 10
		l1.Val = l1.Val % 10

		if l1.Next == nil {
			if carry == 0 {
				l1.Next = l2.Next
				break
			}
			l1.Next = &ListNode{}
		}

		if l2.Next == nil {
			if carry == 0 {
				break
			}
			l2.Next = &ListNode{}
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return reverse(l)
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
