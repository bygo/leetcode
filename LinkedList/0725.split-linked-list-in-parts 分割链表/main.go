package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/split-linked-list-in-parts

func splitListToParts(head *ListNode, k int) []*ListNode {
	cur := head
	l := 0
	// 长度
	for cur != nil {
		cur = cur.Next
		l++
	}

	width := l / k
	rem := l % k
	var res = make([]*ListNode, k)

	for i := 0; i < k; i++ {
		var j = width
		if i < rem { // 前面的比较长
			j++
		}
		res[i] = head
		for 1 < j { // 移动到最后一个
			head = head.Next
			j--
		}
		if head != nil { // 移动到下一个 并且切割
			head, head.Next = head.Next, nil
		}
	}
	return res
}
