package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Title: Linked List Components
//Link: https://leetcode-cn.com/problems/linked-list-components

func numComponents(head *ListNode, G []int) int {
	var m = map[int]struct{}{}

	for _, v := range G {
		m[v] = struct{}{}
	}

	var flag, res int

	for head != nil {
		_, ok := m[head.Val]
		if ok {
			flag++
		} else {
			flag = 0
		}
		if flag == 2 {
			res++
		}
		head = head.Next
	}

	return res
}


