package main

type ListNode struct {
	Val  int
	Prev *ListNode
	Next *ListNode
}

// https://leetcode-cn.com/problems/design-linked-list

type MyLinkedList struct {
	Head *ListNode
	Tail *ListNode
	Size int
}

func Constructor() MyLinkedList {
	head, tail := &ListNode{Val: -1}, &ListNode{Val: -1}
	head.Next = tail
	tail.Prev = head
	return MyLinkedList{
		Head: head,
		Tail: tail,
	}
}

func (this *MyLinkedList) Get(index int) int {
	n := this.get(index)
	if n == nil {
		return -1
	}
	return n.Val
}

func (this *MyLinkedList) get(index int) *ListNode {
	if index <= this.Size/2 {
		return this.getHead(index)
	} else {
		return this.getTail(this.Size - index - 1)
	}
}

func (this *MyLinkedList) getHead(index int) *ListNode {
	cur := this.Head
	for cur.Next != nil && 0 <= index {
		cur = cur.Next
		index--
	}
	return cur
}

func (this *MyLinkedList) getTail(index int) *ListNode {
	cur := this.Tail
	for cur.Prev != nil && 0 <= index {
		cur = cur.Prev
		index--
	}
	return cur
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.Size < index {
		return
	}
	if index < 0 {
		index = 0
	}

	n := &ListNode{Val: val}
	cur := this.get(index)

	var prev, next *ListNode

	prev = cur.Prev
	next = cur

	prev.Next = n
	n.Prev = prev

	next.Prev = n
	n.Next = next

	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if 0 <= index && index < this.Size {
		cur := this.get(index)
		prev := cur.Prev
		next := cur.Next
		prev.Next = next
		next.Prev = prev

		this.Size--
	}
}
