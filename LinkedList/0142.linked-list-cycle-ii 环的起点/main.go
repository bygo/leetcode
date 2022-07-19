package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/linked-list-cycle-ii

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // 相遇时 在环内
			for slow != head { // 如果以环起点，相遇时必定也是环起点
				slow = slow.Next // 因为 slow 入环  fast 相对于 slow 多走了 head->环起点
				head = head.Next // 所以 相遇时 slow head->环起点 = slow->环起点
			}
			return slow
		}
	}
	return nil
}
