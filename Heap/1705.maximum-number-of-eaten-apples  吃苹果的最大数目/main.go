package main

import "container/heap"

// https://leetcode.cn/problems/maximum-number-of-eaten-apples

// ❓ 吃苹果的最大数目

func eatenApples(apples, days []int) int {
	var cntMax int
	h := hp{}
	day := 0

	// 先循环 以避免 未生产 被 先消费
	for day < len(apples) {
		// 已经腐败的 pop
		for 0 < len(h) && h[0].exp <= day {
			heap.Pop(&h)
		}

		// 合法的
		if 0 < apples[day] {
			heap.Push(&h, pair{day + days[day], apples[day]})
		}

		// 吃一个 最不新鲜的
		if 0 < len(h) {
			h[0].cnt--
			if h[0].cnt == 0 {
				heap.Pop(&h)
			}
			cntMax++
		}
		day++
	}

	for 0 < len(h) {
		// 已经腐败的 pop
		for 0 < len(h) && h[0].exp <= day {
			heap.Pop(&h)
		}

		if len(h) == 0 {
			break
		}

		p := heap.Pop(&h).(pair)
		cntMost := min(p.exp-day, p.cnt) // 还能坚持几天
		cntMax += cntMost
		day += cntMost
	}
	return cntMax
}

type pair struct {
	exp,
	cnt int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].exp < h[j].exp
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(pair))
}

func (h *hp) Pop() interface{} {
	top := len(*h) - 1
	p := (*h)[top]
	*h = (*h)[:top]
	return p
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
