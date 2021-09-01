package main

// https://leetcode-cn.com/problems/design-an-ordered-stream

type OrderedStream struct {
	m   []string
	ptr int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		m: make([]string, n+1),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.m[idKey-1] = value
	var res []string
	for 0 < len(this.m[this.ptr]) {
		res = append(res, this.m[this.ptr])
		this.ptr++
	}

	return res
}
