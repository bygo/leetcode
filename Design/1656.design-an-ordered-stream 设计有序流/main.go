package main

// https://leetcode.cn/problems/design-an-ordered-stream

type OrderedStream struct {
	m   []string
	ptr int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		m: make([]string, n+1),
	}
}

func (os *OrderedStream) Insert(idKey int, value string) []string {
	os.m[idKey-1] = value
	var res []string
	for os.m[os.ptr] != "" {
		res = append(res, os.m[os.ptr])
		os.ptr++
	}

	return res
}
