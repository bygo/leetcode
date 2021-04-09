package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Title: Partition List
//Link: https://leetcode-cn.com/problems/partition-list

func partition(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	smallHead, largeHead := small, large
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
