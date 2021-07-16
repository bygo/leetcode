package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/add-two-numbers-ii

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 []int
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	var carry int
	var prev *ListNode
	for 0 < len(s1) || 0 < len(s2) || 0 < carry {
		t1 := len(s1) - 1
		t2 := len(s2) - 1
		var num int
		if -1 < t1 {
			num += s1[t1]
			s1 = s1[:t1]
		}
		if -1 < t2 {
			num += s2[t2]
			s2 = s2[:t2]
		}
		num += carry
		carry = num / 10
		v := num % 10
		cur := &ListNode{Val: v}
		cur.Next = prev
		prev = cur
	}

	return prev
}
