package main

//Link: https://leetcode-cn.com/problems/online-stock-span

type StockSpanner struct {
	arr   []int
	stack []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: []int{},
	}
}

//[100, 80, 60, 70, 60, 75, 85]

func (this *StockSpanner) Next(price int) int {
	top := len(this.stack) - 1
	for -1 < top && this.arr[this.stack[top]] <= price {
		this.stack = this.stack[:top]
		top--
	}

	l := len(this.arr)
	// 起始 index
	var index = 0
	if -1 < top {
		index = this.stack[top]
	}

	this.stack = append(this.stack, l)
	this.arr = append(this.arr, price)
	return l - index
}
