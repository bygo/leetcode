package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Link: https://leetcode-cn.com/problems/linked-list-components

func numComponents(head *ListNode, G []int) int {
	var m = map[int]struct{}{}

	for _, v := range G {
		m[v] = struct{}{}
	}

	var flag, res int

	for head != nil {
		_, ok := m[head.Val]
		if ok { // 出现一次
			flag++
		} else { // 不存在
			flag = 0
		}
		if flag == 1 { // 出现1次 即算一个组件 之后累计的忽略
			res++
		}
		head = head.Next
	}

	return res
}
