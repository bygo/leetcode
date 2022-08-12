package main

import "container/heap"

// https://leetcode.cn/problems/the-k-weakest-rows-in-a-matrix

type pair struct {
	pow int
	idx int
}

type hh []*pair

func (h hh) Len() int { return len(h) }
func (h hh) Less(i, j int) bool {
	return h[i].pow < h[j].pow || h[i].pow == h[j].pow && h[i].idx < h[j].idx
}
func (h hh) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *hh) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

func (h *hh) Push(x interface{}) {
	*h = append(*h, x.(*pair))
}

func kWeakestRows(mat [][]int, k int) []int {
	var h = hh{}
	for i, m := range mat {
		heap.Push(&h, &pair{
			pow: binarySearch(m),
			idx: i,
		})
	}
	var res []int
	for 0 < k {
		k--
		p := heap.Pop(&h).(*pair)
		res = append(res, p.idx)
	}
	return res
}

func binarySearch(m []int) int {
	l, r := 0, len(m)
	for l < r {
		mid := l + (r-l)/2
		if m[mid] == 1 {
			l = mid + 1
		} else if m[mid] == 0 {
			r = mid
		}
	}
	return l - 1
}
