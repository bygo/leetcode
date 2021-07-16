package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list

// 前缀和
func removeZeroSumSublists(head *ListNode) *ListNode {
	// 前缀和
	m := map[int]*ListNode{}
	zero := &ListNode{Next: head}

	// 保存最后一个 sum 出现的位置
	cur, sum := zero, 0
	for cur != nil {
		sum += cur.Val
		m[sum] = cur
		cur = cur.Next
	}

	// 出现 相同的值  删除所有节点
	cur, sum = zero, 0
	for cur != nil {
		sum += cur.Val
		cur.Next = m[sum].Next // 只有一个节点 用自己的节点
		cur = cur.Next
	}
	return zero.Next
}

// 暴力法
func removeZeroSumSublists(head *ListNode) *ListNode {
	zero := &ListNode{Next: head}
	pre := zero
	cur := pre.Next
	for pre.Next != nil {
		sum := 0
		for cur != nil {
			sum += cur.Val
			if sum == 0 {
				break
			}
			cur = cur.Next
		}
		if sum == 0 { // 循环完 = 0 删除区间
			cur = cur.Next
			pre.Next = cur
		} else { // 进入下一次循环
			pre = pre.Next
			cur = pre.Next
		}
	}

	return zero.Next
}
