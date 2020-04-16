package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l := head      //左指针
	r := head.Next //右指针
	for r != nil {
		tmp := r.Next //临时
		r.Next = l    //右指针.next 指向左指针
		l = r         //左指针移动
		r = tmp       //右指针移动
	}
	head.Next = nil
	return l
}
