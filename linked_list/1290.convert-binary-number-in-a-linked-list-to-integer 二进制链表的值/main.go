package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//Title: Convert Binary Number in a Linked List to Integer
// https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer

func getDecimalValue(head *ListNode) int {
	var i int
	for head != nil {
		i = i<<1 + head.Val
		head = head.Next
	}
	return i
}
