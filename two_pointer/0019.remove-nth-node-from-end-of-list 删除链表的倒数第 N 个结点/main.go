package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list

// zero->1->2->3->nil
// 因 left 与 right 相距 n
// 当 right = nil
// left n+1
// left.Next = n
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	zero := &ListNode{Next: head}
	l, r := zero, zero
	// r 先走 n 步
	for -1 < n {
		r = r.Next
		n--
	}
	// r 走到末尾
	for r != nil {
		l = l.Next
		r = r.Next
	}
	l.Next = l.Next.Next
	return zero.Next
}
