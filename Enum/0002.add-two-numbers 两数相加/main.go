package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/add-two-numbers

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	head := l1
	for {
		l1.Val += carry + l2.Val
		carry = l1.Val / 10
		l1.Val %= 10

		// 其一链表Next为空时，尝试链表合并
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

	return head
}
