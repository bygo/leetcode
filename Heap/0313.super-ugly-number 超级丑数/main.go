package main

import (
	"container/heap"
	"sort"
)

// https://leetcode-cn.com/problems/super-ugly-number

type h struct {
	sort.IntSlice
}

func (h *h) Pop() interface{} {
	n := len(h.IntSlice) - 1
	x := h.IntSlice[n]
	h.IntSlice = h.IntSlice[:n]
	return x
}

func (h *h) Push(i interface{}) {
	x := i.(int)
	h.IntSlice = append(h.IntSlice, x)
}

func nthSuperUglyNumber(n int, primes []int) int {
	var res int
	m := map[int]struct{}{1: {}}
	h := &h{[]int{1}}
	for i := 0; i < n; i++ {
		res = heap.Pop(h).(int)
		for _, prime := range primes {
			v := res * prime
			_, ok := m[v]
			if !ok {
				m[v] = struct{}{}
				heap.Push(h, v)
			}
		}
	}
	return res
}
