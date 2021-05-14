package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Link: https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list


// zero->1->2->3->nil
// 因 left 与 right 相距 n
// 当 right = nil
// left n+1
// left.Next = n
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	zero := &ListNode{Next: head}
	left, right := zero, zero
	for 0 < n {
		right = right.Next
		n--
	}
	for right != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return zero.Next
}
