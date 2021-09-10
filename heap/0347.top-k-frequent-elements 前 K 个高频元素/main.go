package main

import "container/heap"

// https://leetcode-cn.com/problems/top-k-frequent-elements

type h []*e

type e struct {
	cnt int
	ele int
}

func (h h) Len() int           { return len(h) }
func (h h) Less(i, j int) bool { return h[j].cnt < h[i].cnt }
func (h h) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *h) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

func (h *h) Push(x interface{}) {
	*h = append(*h, x.(*e))
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for i := range nums {
		m[nums[i]]++
	}

	var h = &h{}
	for i := range m {
		heap.Push(h, &e{
			cnt: m[i],
			ele: i,
		})
	}

	var res []int
	for 0 < k {
		res = append(res, heap.Pop(h).(*e).ele)
		k--
	}
	return res
}
