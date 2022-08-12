package main

// https://leetcode.cn/problems/maximum-frequency-stack

type FreqStack struct {
	max   int
	freq  map[int]int
	group map[int][]int
}

func Constructor() FreqStack {
	return FreqStack{
		max:   0,
		freq:  map[int]int{},   // 计算频率
		group: map[int][]int{}, // 频率栈
	}
}

func (this *FreqStack) Push(val int) {
	this.freq[val] += 1 // 频率计算
	f := this.freq[val]
	if this.max < f { //
		this.max = f
	}
	if this.group[f] == nil {
		this.group[f] = []int{}
	}
	this.group[f] = append(this.group[f], val)
}

func (this *FreqStack) Pop() int {
	top := len(this.group[this.max]) - 1
	val := this.group[this.max][top]
	this.freq[val] -= 1 //频率计算

	this.group[this.max] = this.group[this.max][:top]
	if 0 == top {
		this.max--
	}
	return val
}
