package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Link: https://leetcode-cn.com/problems/merge-k-sorted-lists

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if 0 == l {
		return nil
	}
	for 1 < l {
		for i := 0; i < l/2; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[l-i-1])
		}
		l = (l + 1) / 2
	}
	return lists[0]
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := &ListNode{}
	res := prev
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 == nil {
		prev.Next = l2
	} else if l2 == nil {
		prev.Next = l1
	}
	return res.Next
}
