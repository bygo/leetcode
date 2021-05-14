package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Link: https://leetcode-cn.com/problems/odd-even-linked-list

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	even := head.Next
	odd := head
	evenHead := even
	for even != nil && even.Next != nil { // 如果还有奇数
		odd.Next = even.Next // 1.Next = 2.Next
		odd = odd.Next       // 1 = 3
		even.Next = odd.Next // 2.Next = 3.Next
		even = even.Next     // 2 = 4
	}
	odd.Next = evenHead
	return head
}
