package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Title: Odd Even Linked List
//Link: https://leetcode-cn.com/problems/odd-even-linked-list

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func partition(head *ListNode, x int) *ListNode {
	var small, large = &ListNode{}, &ListNode{}
	var smallHead, largeHead = small, large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}

	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}

