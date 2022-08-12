package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/design-front-middle-back-queue

type FrontMiddleBackQueue struct {
	d []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{d: []int{}}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.d = append(this.d, 0)
	l := len(this.d)
	copy(this.d[1:], this.d[0:l-1])
	this.d[0] = val
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	m := len(this.d) / 2
	this.d = append(this.d, 0)
	copy(this.d[m+1:], this.d[m:])
	this.d[m] = val
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.d = append(this.d, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	top := len(this.d) - 1
	if top < 0 {
		return -1
	}
	v := this.d[0]
	this.d = this.d[1:]
	return v
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	top := len(this.d) - 1
	if top < 0 {
		return -1
	}
	m := top / 2
	v := this.d[m]

	copy(this.d[m:top], this.d[m+1:])
	this.d = this.d[:top]
	return v
}

func (this *FrontMiddleBackQueue) PopBack() int {
	top := len(this.d) - 1
	if top < 0 {
		return -1
	}
	v := this.d[top]
	this.d = this.d[:top]
	return v
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
