package main

type ListNode struct {
	Coefficient int
	Power       int
	Next        *ListNode
}

// https://leetcode.cn/problems/add-two-polynomials-represented-as-linked-lists

func addPoly(lh, rh *ListNode) *ListNode {
	l := &ListNode{Next: lh}
	r := &ListNode{Next: rh}

	for l.Next != nil && r.Next != nil {
		if l.Next.Power == r.Next.Power {
			l.Next.Coefficient += r.Next.Coefficient
			if l.Next.Coefficient == 0 {
				l.Next = l.Next.Next
			}
			l = l.Next
			r = r.Next
		} else if l.Next.Power < r.Next.Power {
			next := l.Next
			l.Next = r.Next
			r = r.Next
			l.Next.Next = next
		} else if r.Next.Power < l.Next.Power {
			next := r.Next
			r.Next = l.Next
			l = l.Next
			r.Next.Next = next
		}
	}
	return lh
}
