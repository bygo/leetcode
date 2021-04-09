package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Title: Remove Nth Node From End of List
//Link: https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	zero := &ListNode{
		Next: head,
	}
	left, right := zero, zero
	for -1 < n {
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

/**
思路:
1.双指针
2.头部加入空节点
*/
