package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Title: Reverse Nodes in k-Group
//Link: https://leetcode-cn.com/problems/reverse-nodes-in-k-group

func reverseKGroup(head *ListNode, k int) *ListNode {
	empty := &ListNode{}
	empty.Next = head
	prev := empty
	for head != nil {
		for i := 0; i < k-1 && head != nil; i++ {
			head = head.Next
		}
		if head == nil {
			break
		}
		first := prev.Next
		next := head.Next
		head.Next = nil
		prev.Next = reverse(first)
		first.Next = next
		head = next
		prev = first
	}

	return empty.Next
}

func reverse(first *ListNode) *ListNode {
	if first == nil || first.Next == nil {
		return first
	}
	l := first
	r := first.Next
	for r != nil {
		tmp := r.Next
		r.Next = l
		l = r
		r = tmp
	}
	first.Next = nil
	return l
}



/**

 */
