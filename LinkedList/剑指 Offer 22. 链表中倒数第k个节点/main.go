package main

// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	s, f := head, head
	for 0 < k {
		f = f.Next
		k--
	}

	for f != nil {
		s = s.Next
		f = f.Next
	}
	return s
}
