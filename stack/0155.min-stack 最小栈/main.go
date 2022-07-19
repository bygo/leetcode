package main

// https://leetcode.cn/problems/min-stack

type MinStack struct {
	minStack []int
	stack    []int
}

// 每个栈顶的值的最小值 存起来
func Constructor() MinStack {
	return MinStack{
		minStack: []int{1<<63 - 1},
	}
}

func (this *MinStack) Push(val int) {
	top := len(this.minStack) - 1
	this.minStack = append(this.minStack, min(this.minStack[top], val))
	this.stack = append(this.stack, val) //
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
