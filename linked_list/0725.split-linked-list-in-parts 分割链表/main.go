package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/split-linked-list-in-parts

func splitListToParts(root *ListNode, k int) []*ListNode {
	cur := root
	l := 0
	// 长度
	for cur != nil {
		cur = cur.Next
		l++
	}

	width := l / k
	rem := l % k
	var res []*ListNode

	for i := 0; i < k; i++ {
		var j = width
		if i < rem { // 前面的比较长
			j += 1
		}
		res = append(res, root)
		for 1 < j { // 只有一个 原地不动
			root = root.Next
			j--
		}
		if root != nil { // 移动到下一个 并且切割
			next := root.Next
			root.Next = nil
			root = next
		}
	}
	return res
}
